package query

import (
	"bufio"
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"strings"

	"math"
	"os"
	"sync"
	"sync/atomic"
)

//SliceFetchAndExport multi-goroutine 写入缓冲channel 一个消费者消费并写文件
func SliceFetchAndExport(ctx *basic.ExeElasticSQLCtx, indices []string, export *grammer.ExportClause, searchSource *elastic.SearchSource, field string, sliceMax int, limit int64)  (int64,error) {
	// log(10, size)
	if sliceMax <= 1 || sliceMax > ctx.Conf.Query.ScrollSliceMax {
		num := 1
		if limit > 1 {
			num = int(limit)
		}
		threadNum := 1
		if num < ctx.Conf.Query.ScrollSliceThreshold {
			threadNum = 1
		} else {
			threadNum = int(math.Min(math.Floor(math.Log10(5.0*float64(num)/float64(ctx.Conf.Query.ScrollSliceThreshold))+2), float64(ctx.Conf.Query.ScrollSliceMax)))
		}
		sliceMax = threadNum
	}
	client := ctx.Conn.Client

	queue := make(chan *elastic.SearchHit, 500)
	produceEndSignal := make(chan bool)
	consumerEndSignal := make(chan bool)
	var count int64

	var fetch int64

	// 启动消费者
	go func() {
		fetch,_= ConsumeAndExport(ctx,export, queue, produceEndSignal, consumerEndSignal,limit)
	}()

	wg := &sync.WaitGroup{}
	//errorMsg := ""
	// 生产者
	for i := 0; i < sliceMax; i++ {
		wg.Add(1)
		go func() {
			sliceQuery := elastic.NewSliceQuery().Id(i).Max(sliceMax)
			if field != "" {
				sliceQuery.Field(field)
			}
			subScrollService := elastic.NewScrollService(client)
			subScrollService.Slice(sliceQuery).Scroll("2m").Size(ctx.Conf.Query.ScrollSize)
			subScrollService.SearchSource(searchSource)
			subScrollService.Index(indices...)
			response, err := subScrollService.Do(ctx.GCtx)
			if err != nil {
				//errorMsg = "slice scroll do fail"
				return
			}
			hits := response.Hits.Hits
			scrollId := response.ScrollId
			for true {

				len := len(hits)
				if len == 0 {
					break
				}
				atomic.AddInt64(&count, response.Hits.TotalHits.Value)
				for _, hit := range hits {
					// 写入队列
					queue <- hit
				}
				if count >= limit {
					// 不需要继续获取数据
					break
				}
				// 下一轮 retrieve the next batch of results
				scroll := elastic.NewScrollService(client)
				scroll.ScrollId(scrollId)
				scroll.Scroll("1m")
				response, err = scroll.Do(ctx.GCtx)
				if err != nil {
					//errorMsg = "slice scroll do fail"
					return
				}
				hits = response.Hits.Hits
				scrollId = response.ScrollId
			}

			// clear scroll
			client.ClearScroll(scrollId).Do(ctx.GCtx)

			wg.Done()
		}()
	}

	// 主线程等待生产者结束
	wg.Wait()
	// 生产结束发送信号
	produceEndSignal <- true

	// 等待消费结束
	<-consumerEndSignal
	// 消费结束

	return fetch,nil

}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//ConsumeAndExport 作为消费者
func ConsumeAndExport(ctx *basic.ExeElasticSQLCtx, export *grammer.ExportClause, queue <-chan *elastic.SearchHit, end <-chan bool, consumerEnd chan<- bool , limit int64) (int64,error){
	var (
		file *os.File
		err  error
	)
	dir:= ctx.Conf.DataDir
	if strings.HasSuffix(dir,"/"){
		dir = dir[0:len(dir)-1]
	}
	exportFilePath := dir+"/" + export.FileName
	if Exists(exportFilePath) {
		//使用追加模式打开文件
		file, err = os.OpenFile(exportFilePath, os.O_APPEND, 0666)
		if err != nil {
			return 0,errors.Wrap(err, "open file "+exportFilePath)
		}
	} else {
		file, err = os.Create(exportFilePath) //创建文件
		if err != nil {
			return 0,errors.Wrap(err, "create file "+exportFilePath)
		}
	}
	defer file.Close()
	defer func() { consumerEnd <- true }()

	writer := bufio.NewWriter(file)

	var mapper func(hit *elastic.SearchHit) []byte
	if export.FileType == grammer.EXPORT_CSV {
		mapper = GetCsvMapper(export.FetchCode, export.Fields)
	} else {
		mapper = GetJSONMapper(export.FetchCode)
	}
	var fetch int64

	for {
		select {
		case x := <-queue:
			line := mapper(x)
			writer.Write(line)
			writer.WriteByte('\n')
			fetch ++

		case flag := <-end: // 终止flag
			fmt.Println("flag=", flag)
			writer.Flush()
			return  fetch,nil//不能用break，因为只跳出select这一层
		case <-ctx.GCtx.Done():
			return fetch,ctx.GCtx.Err()

		}
	}
}
