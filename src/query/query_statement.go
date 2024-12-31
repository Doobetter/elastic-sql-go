/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package query

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Doobetter/elastic-sql-go/src/basic"
	"github.com/Doobetter/elastic-sql-go/src/common"
	"github.com/Doobetter/elastic-sql-go/src/grammer"
	"github.com/olivere/elastic/v7"
	"log"
	"strings"
)

// QueryStatement
//总条数在 scroll_slice_threshold 以下的都可以做后续join 等操作
//总条数在 scroll_slice_threshold 以上的不可以join、mem-sort,只能export到文件

type QueryStatement struct {
	basic.ProcessUnit
	Front      string   // 前置Statement名字
	Indexes    []string // 索引名
	Fields     []string // 字段名
	Highlight  *grammer.HighlightAdapter
	ExprFields map[string]*grammer.ScriptAdapter

	// for scroll
	ScrollKeepTime string // e.g. "2m" or "30s"
	ScrollId       string

	MinScore float64

	Where    grammer.Expression
	Rescore  *grammer.RescoreAdapter
	Collapse *grammer.CollapseAdapter

	// sort clause
	SortBys []*grammer.SortAdapter
	MemSort []*grammer.SortAdapter

	// limit clause
	From int
	Size int

	SliceMax   int
	SliceFiled string

	MapName string
	Export  *grammer.ExportClause

	fetchAllSource bool     // 对应查询语句中有 '*'
	fetchSource    []string // 通过_source 获取的字段
	resultSchema   []string // 结果中需要的字段包括 meta中和script、_source中字段
	fetchCode      int      // 初始为0，之后加上需要获取的code

	SearchSource  *elastic.SearchSource
	searchService *elastic.SearchService
	scrollService *elastic.ScrollService
}

func NewQueryStatement() *QueryStatement {
	stat := new(QueryStatement)
	stat.ProcessUnit.Processes = map[int]basic.Void{}
	return stat
}

//// 实现 覆盖 接口、父类 函数开始

func (s *QueryStatement) Init(ctx *basic.ExeElasticSQLCtx) error {
	if s.ScrollId != "" || s.ScrollKeepTime != "" {
		//s.searchRequest.Scroll(s.ScrollKeepTime)
		s.scrollService = elastic.NewScrollService(ctx.Conn.Client)
		// TODO scroll
	} else {
		// by query
		s.SearchSource = elastic.NewSearchSource()
		// request level
		s.searchService = elastic.NewSearchService(ctx.Conn.Client).SearchSource(s.SearchSource)
		s.searchService.Index(s.Indexes...) //todo 对带*的index做特殊处理
		s.searchService.AllowNoIndices(true)
		// search source level
		s.SearchSource.TrackTotalHits(true)
		s.SetFetchSourceAndSchema()
		if s.Where != nil {
			s.SearchSource.Query(grammer.AdaptToQueryBuilder(s.Where, ctx))
			if s.Rescore != nil {
				// todo
			}
		}
		if s.Collapse != nil {
			// todo
		}

		if s.MinScore > 0 {
			s.SearchSource.MinScore(s.MinScore)
		}

		// 排序设置
		for _, sortBy := range s.SortBys {
			s.SearchSource.SortBy(sortBy.ToFieldSortBuilder())
		}

		// 获取条数设置
		if s.From > 0 {
			s.SearchSource.From(s.From)
		}
		if s.Size <= 0 {
			s.SearchSource.Size(ctx.Conf.Query.DefaultSize)
		} else if s.Size > 0 && s.Size < 10000 {
			s.SearchSource.Size(s.Size)
		} else {
			// 没有设置size 使用默认的配置
			s.SearchSource.Size(ctx.Conf.Query.DefaultSize)
		}
		if s.IsViaNON() {

		} else {
			s.SearchSource.Size(ctx.Conf.Query.DefaultSize)
		}

	}
	if s.MapName != "" {
		s.Name = s.MapName
	}

	return nil
}

func (s *QueryStatement) Execute(ctx *basic.ExeElasticSQLCtx) error {
	resultSet := basic.NewResultSet(s.Name)
	ctx.AddResultSet(s.Name, resultSet)
	searchResult, err := s.getRealSearchResult(ctx.GCtx)
	if err != nil {
		resultSet.ErrMsg = err.Error()
		return err
	}
	resultSet.SearchResponse = searchResult
	resultSet.Schemas = s.resultSchema
	total := searchResult.TotalHits()
	resultSet.Total = total
	if s.Size > 0 {
		resultSet.Count = common.MinInt64(total, int64(s.Size))
	} else {
		resultSet.Count = total
	}

	if s.IsViaNON() {
		// 不导出文件、无其他后续处理依赖
		resultSet.Data = s.ParseSearchResultData(searchResult)
		resultSet.FetchSize = int64(len(resultSet.Data))
		resultSet.ScrollId = searchResult.ScrollId
		// todo mem-sort
	} else {
		if s.Size > 0 { // limit from,size
			if s.Size < 10000 || total < 10000 {
				s.fetch(resultSet, searchResult)
				// todo mem-sort
			} else {
				if s.Size < ctx.Conf.Query.ScrollSliceThreshold {
					// normal scroll get
					s.normalScrollFetch(ctx, resultSet)
				} else {
					// big-big-file
					s.sliceScrollFetch(ctx, resultSet)
				}
			}

		} else {
			if total < 10000 {
				s.fetch(resultSet, searchResult)
				// todo mem-sort
			} else {
				if total < int64(ctx.Conf.Query.ScrollSliceThreshold) {
					// normal scroll get
					s.normalScrollFetch(ctx, resultSet)
				} else {
					s.sliceScrollFetch(ctx, resultSet)

				}
			}
		}
	}
	return nil
}

func (s *QueryStatement) Clean(ctx *basic.ExeElasticSQLCtx) {

}
func (s *QueryStatement) GetExportFileName() string {
	if s.Export != nil {
		return s.Export.FileName
	}
	return ""
}

func (s *QueryStatement) DefaultNameSuffix() string {
	return "_query"
}

// getRealSearchResult 区分 scroll 与普通查询
func (s *QueryStatement) getRealSearchResult(ctx context.Context) (*elastic.SearchResult, error) {
	if s.searchService != nil {
		log.Printf("DSL POST %s/_search %s", strings.Join(s.Indexes, ","), common.JSONPrettyStrWithErr(s.SearchSource.Source()))
		return s.searchService.Do(ctx)
	} else {
		//scroll
		return s.scrollService.Do(ctx)
	}
}

func (s *QueryStatement) SetFetchSourceAndSchema() {
	for _, f := range s.Fields {
		if "*" == s.Fields[0] {
			s.fetchAllSource = true
			s.fetchCode |= _allSource
			// TODO GET all fields by mapping
		} else if strings.HasPrefix(f, "_") {
			if f == "_index" {
				s.fetchCode |= _index
			} else if f == "_id" {
				s.fetchCode |= _id
			} else if f == "_score" {
				s.fetchCode |= _score
			}
			s.resultSchema = append(s.resultSchema, f)
		} else {
			s.fetchCode |= _someSource
			s.fetchSource = append(s.fetchSource, f)
			s.resultSchema = append(s.resultSchema, f)
		}
	}
	if len(s.fetchSource) > 0 {
		s.SearchSource.FetchSourceIncludeExclude(s.fetchSource, nil)
	} else if s.fetchAllSource == false {
		// 不获取字段只获取doc的元数据
		s.SearchSource.FetchSource(false)
	}

	// for 高亮
	if s.Highlight != nil {
		s.fetchCode |= _hilight
		s.SearchSource.Highlight(s.Highlight.ToHighlightBuilder())
		s.resultSchema = append(s.resultSchema, s.Highlight.FieldSchema...)
	}
	// for 脚本生成字段
	//TODO
	if s.Export != nil {
		s.Export.FetchCode = s.fetchCode
		s.Export.ResetFields(s.resultSchema)
	}

}

// ParseSearchResultData 将SearchResult解析到 []map
func (s *QueryStatement) ParseSearchResultData(sr *elastic.SearchResult) []map[string]interface{} {
	var data []map[string]interface{}
	if sr == nil {
		return data
	}
	hits := sr.Hits.Hits
	// _id 比较特殊需要速度快
	if s.fetchCode == _id {
		// just _id
		for _, hit := range hits {
			row := make(map[string]interface{}, 1)
			row["_id"] = hit.Id
			data = append(data, row)
		}

	} else {
		for _, hit := range hits {
			row := s.getSourceAndHighlights(hit)
			if s.fetchCode&_index != 0 {
				row["_index"] = hit.Index
			}
			if s.fetchCode&_id != 0 {
				row["_id"] = hit.Id
			}
			if s.fetchCode&_score != 0 {
				row["_score"] = hit.Score
			}
			data = append(data, row)
		}
	}
	return data
}

func (s *QueryStatement) getSourceAndHighlights(hit *elastic.SearchHit) map[string]interface{} {
	row := make(map[string]interface{})
	if hit.Source != nil {
		json.Unmarshal(hit.Source, &row)
	}
	if hit.Highlight != nil {
		for f, vArray := range hit.Highlight {
			as := s.Highlight.FieldAndSchema[f]
			row[as] = vArray[0] // 只获一条
		}
	}
	return row
}

func (s *QueryStatement) fetch(rs *basic.ResultSet, sr *elastic.SearchResult) {
	if s.IsViaJOIN() {
		rs.Data = s.ParseSearchResultData(sr)
		rs.FetchSize = int64(len(rs.Data))
	}
	if s.IsViaExport() {
		if s.Export.FileType == grammer.EXPORT_CSV {
			rs.Headers = s.Export.Headers
			rs.DataStr = s.CSV(sr)
		} else {
			//导出JSON格式数据
			rs.DataStr = s.JSON(sr)
		}
	}
}

func (s *QueryStatement) normalScrollFetch(ctx *basic.ExeElasticSQLCtx, resultSet *basic.ResultSet) {
	if s.IsViaExport() {
		fetched, err := ScrollFetchAndExport(ctx, s.Indexes, s.Export, s.SearchSource, int64(s.Size))
		resultSet.FetchSize = fetched
		if err != nil {
			resultSet.ErrMsg = err.Error()
			return
		}
		resultSet.WarnMsg = "数据量超过scroll_slice_threshold直接写到文件"
	} else if s.IsViaJOIN() {
		// 将json格式数据放到resultSet中
		ScrollFetch(ctx, s.fetchCode, resultSet, s.Indexes, s.SearchSource, int64(s.Size))
	}

}

func (s *QueryStatement) sliceScrollFetch(ctx *basic.ExeElasticSQLCtx, resultSet *basic.ResultSet) {
	if s.IsJustViaExport() {
		fetched, err := SliceFetchAndExport(ctx, s.Indexes, s.Export, s.SearchSource, s.SliceFiled, s.SliceMax, int64(s.Size))
		resultSet.FetchSize = fetched
		if err != nil {
			resultSet.ErrMsg = err.Error()
			return
		}
		resultSet.WarnMsg = "数据量超过scroll_slice_threshold直接写到文件"

	} else {
		resultSet.ErrMsg = "数据量超过scroll_slice_threshold只支持导出到文件，不支持join等后续操作"
	}
}

func (s *QueryStatement) ToLine(hit *elastic.SearchHit) string {
	var line bytes.Buffer
	row := s.getSourceAndHighlights(hit)
	for _, f := range s.resultSchema {
		if "_id" == f {
			line.WriteString(hit.Id)
		} else if "_index" == f {
			line.WriteString(hit.Index)
		} else if "_score" == f {
			line.WriteString(common.Float64ToString(*hit.Score))
		} else {
			val := row[f]
			if val == nil {
				line.WriteString("")
			}
			//todo
			line.WriteString(fmt.Sprintf("%v", val))
			//switch val.(type) {
			//case string,float32,float64,int,int64:
			//
			//case []interface{}:
			//	// to
			//	line.WriteString("--")
			//
			//}
		}
	}
	return line.String()
}
func (s *QueryStatement) CSV(sr *elastic.SearchResult) string {
	var data bytes.Buffer

	hits := sr.Hits.Hits
	// _id 比较特殊需要速度快
	if s.fetchCode == _id {
		// just _id
		for _, hit := range hits {
			data.WriteString(hit.Id)
			data.WriteByte('\n')
		}

	} else {
		for _, hit := range hits {
			row := s.getSourceAndHighlights(hit)
			if s.fetchCode&_index != 0 {
				row["_index"] = hit.Index
			}
			if s.fetchCode&_id != 0 {
				row["_id"] = hit.Id
			}
			if s.fetchCode&_score != 0 {
				row["_score"] = hit.Score
			}
			bs, _ := json.Marshal(row)
			data.Write(bs)
			data.WriteByte('\n')
		}
	}
	return data.String()
}

func (s *QueryStatement) JSON(sr *elastic.SearchResult) string {
	var data bytes.Buffer

	hits := sr.Hits.Hits
	// _id 比较特殊需要速度快
	if s.fetchCode == _id {
		// just _id
		for _, hit := range hits {
			row := make(map[string]interface{}, 1)
			row["_id"] = hit.Id
			bs, _ := json.Marshal(row)
			data.Write(bs)
			data.WriteByte('\n')
		}

	} else {
		for _, hit := range hits {
			row := s.getSourceAndHighlights(hit)
			if s.fetchCode&_index != 0 {
				row["_index"] = hit.Index
			}
			if s.fetchCode&_id != 0 {
				row["_id"] = hit.Id
			}
			if s.fetchCode&_score != 0 {
				row["_score"] = hit.Score
			}
			bs, _ := json.Marshal(row)
			data.Write(bs)
			data.WriteByte('\n')
		}
	}
	return data.String()
}
