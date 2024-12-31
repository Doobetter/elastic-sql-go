package grammer

import (
	"github.com/Doobetter/elastic-sql-go/src/conf"
	"fmt"
	"os"
)

const (
	EXPORT_JSON = "JSON"
	EXPORT_CSV  = "CSV"
)

//ExportClause 导出文件
type ExportClause struct {
	Fields     []string
	FileName   string
	FileType   string
	Sep        string
	SaveHeader bool
	Headers    []string
	FetchCode  int
}

func NewExportClause() *ExportClause {
	return &ExportClause{
		Fields:     nil,
		FileName:   "",
		FileType:   EXPORT_CSV, // 默认
		Sep:        conf.CSV_FIELD_SEP,
		SaveHeader: false,
		Headers:    nil,
	}
}
//ResetFields 如果Fields没有设置 就重置
func (c* ExportClause) ResetFields(schema []string)  {
	if len(c.Fields) <=0{
		c.Fields = schema
	}
}
// 判断所给路径文件/文件夹是否存在
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

// 消费者
func (c *ExportClause) export() {
	var (
		file *os.File
		err  error
	)
	if Exists(c.FileName) {
		//使用追加模式打开文件
		file, err = os.OpenFile(c.FileName, os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("Open file err =", err)
			return
		}
	} else {
		file, err = os.Create(c.FileName) //创建文件
		if err != nil {
			fmt.Println("file create fail")
			return
		}
	}
	defer file.Close()
	//writer := bufio.NewWriter(file)
}
