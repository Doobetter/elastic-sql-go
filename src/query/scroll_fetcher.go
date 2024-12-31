/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package query

import (
	"context"
	"errors"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/olivere/elastic/v7"

	"io"
	"log"

	"golang.org/x/sync/errgroup"
)

// ScrollFetchAndExport one-goroutine 写入缓冲channel 一个消费者消费并写文件
func ScrollFetchAndExport(ctx *basic.ExeElasticSQLCtx, indices []string, export *grammer.ExportClause, searchSource *elastic.SearchSource, limit int64) (int64, error) {

	client := ctx.Conn.Client

	queue := make(chan *elastic.SearchHit, 200)
	produceEndSignal := make(chan bool)
	consumerEndSignal := make(chan bool)

	var count int64 // 生产条数

	var fetch int64 // 消费条数

	errGo, _ := errgroup.WithContext(context.Background())
	// 启动消费者
	errGo.Go(func() error {
		var err error
		fetch, err = ConsumeAndExport(ctx, export, queue, produceEndSignal, consumerEndSignal, limit)
		return err
	})
	// 生产者
	errGo.Go(func() error {
		defer close(queue)
		scroll := elastic.NewScrollService(client)
		scroll.Scroll("2m").Size(ctx.Conf.Query.ScrollSize)
		scroll.Index(indices...).SearchSource(searchSource)
		var scrollId string
		for true {
			response, err := scroll.Do(ctx.GCtx)
			if err == io.EOF {
				scrollId = response.ScrollId
				return nil // all results retrieved

			}
			if err != nil {
				return err
			}
			hits := response.Hits.Hits

			len := len(hits)
			if len == 0 {
				scrollId = response.ScrollId
				break
			}
			count += int64(len)
			for _, hit := range hits {
				select {
				case queue <- hit:
					// 写入队列
				case <-ctx.GCtx.Done():
					return ctx.GCtx.Err()
				}
			}
			if count >= limit {
				scrollId = response.ScrollId
				// 不需要继续获取数据
				break
			}
			// 下一轮 retrieve the next batch of results

		}
		// 生产结束发送信号
		produceEndSignal <- true
		// clear scroll
		if scrollId != "" {
			client.ClearScroll(scrollId).Do(ctx.GCtx)
		}
		return nil
	})

	// 主线程等待 等待消费结束
	<-consumerEndSignal
	// 消费结束
	if fetch != count {
		errors.New("写入文件结果数与获取数不等")
	}
	log.Println("write=", fetch, ", read=", count)
	// Check whether any goroutines failed.
	if err := errGo.Wait(); err != nil {
		return fetch, err
	}
	return fetch, nil

}

func ScrollFetch(ctx *basic.ExeElasticSQLCtx, fetchCode int, resultSet *basic.ResultSet, indices []string, searchSource *elastic.SearchSource, limit int64) error {

	mapper := GetMapMapper(fetchCode)
	var count int64 // 生产条数
	scroll := elastic.NewScrollService(ctx.Conn.Client)
	scroll.Scroll("2m").Size(ctx.Conf.Query.ScrollSize)
	scroll.Index(indices...).SearchSource(searchSource)
	var scrollId string
	for true {
		response, err := scroll.Do(ctx.GCtx)
		if err == io.EOF {
			scrollId = response.ScrollId
		}
		if err != nil {
			return err
		}
		hits := response.Hits.Hits

		len := len(hits)
		if len == 0 {
			scrollId = response.ScrollId
			break
		}
		count += int64(len)
		for _, hit := range hits {
			select {
			// 写入队列
			case <-ctx.GCtx.Done():
				return ctx.GCtx.Err()
			default:
				resultSet.Data = append(resultSet.Data, mapper(hit))
			}
		}
		if count >= limit {
			scrollId = response.ScrollId
			// 不需要继续获取数据
			break
		}
		// 下一轮 retrieve the next batch of results

	}

	// clear scroll
	if scrollId != "" {
		ctx.Conn.Client.ClearScroll(scrollId).Do(ctx.GCtx)
	}

	return nil

}
