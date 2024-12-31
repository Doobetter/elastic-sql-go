/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package query

import (
	"bytes"
	"encoding/json"
	"github.com/Doobetter/elastic-sql-go/src/conf"
	"github.com/olivere/elastic/v7"
	"strconv"
)

func GetCsvMapper(fetchCode int, fields []string) func(hit *elastic.SearchHit) []byte {
	var mapper func(hit *elastic.SearchHit) []byte
	// 只有_id
	if fetchCode == _id {
		mapper = func(hit *elastic.SearchHit) []byte {
			return []byte(hit.Id)
		}
	} else if fetchCode == _allSource {
		// 只需要获取 不需要考虑顺序
		mapper = func(hit *elastic.SearchHit) []byte {

			row := make(map[string]interface{})

			json.Unmarshal(hit.Source, &row)

			if len(row) > 0 {
				bs := bytes.Buffer{}

				for _, v := range row {
					bs.WriteString(parseSourceValueToString(v))
					bs.WriteString(conf.CSV_FIELD_SEP)
				}
				if bs.Len() > 0 {
					bs.UnreadByte()
				}
				bs.WriteByte('\n')

				return bs.Bytes()
			} else {
				return nil // todo test
			}
		}
	} else if fetchCode == _someSource {
		// 需要考虑顺序
		mapper = func(hit *elastic.SearchHit) []byte {

			row := make(map[string]interface{})

			json.Unmarshal(hit.Source, &row)

			if len(row) > 0 {
				bs := bytes.Buffer{}
				for _, field := range fields {
					if v, ok := row[field]; ok {
						bs.WriteString(parseSourceValueToString(v))
					} else {
						bs.WriteString("")
					}
					bs.WriteString(conf.CSV_FIELD_SEP)
				}
				if bs.Len() > 0 {
					bs.UnreadByte()
				}
				bs.WriteByte('\n')

				return bs.Bytes()
			} else {
				return nil // todo test
			}
		}
	} else {
		// 含有 _meta 和 source 的混合
		mapper = func(hit *elastic.SearchHit) []byte {

			row := make(map[string]interface{})

			json.Unmarshal(hit.Source, &row)
			bs := bytes.Buffer{}
			for _, field := range fields {
				if field == "_index" {
					bs.WriteString(hit.Index)
				} else if field == "_id" {
					bs.WriteString(hit.Id)
				} else if field == "_score" {
					bs.WriteString(strconv.FormatFloat(*hit.Score, 'f', 32, 64))
				} else if v, ok := row[field]; ok {
					bs.WriteString(parseSourceValueToString(v))
				} else {
					bs.WriteString("")
				}
				bs.WriteString(conf.CSV_FIELD_SEP)
			}
			if bs.Len() > 0 {
				bs.UnreadByte()
			}
			bs.WriteByte('\n')

			return bs.Bytes()
		}
	}

	return mapper
}

func parseSourceValueToString(v interface{}) string {
	switch v.(type) {
	case string:
		return v.(string)
	case int64:
		return strconv.FormatInt(v.(int64), 64)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v.(bool))
	default:
		// map or array ignore
		return ""
	}
}

func GetJSONMapper(fetchCode int) func(hit *elastic.SearchHit) []byte {
	var jsonMapper func(hit *elastic.SearchHit) []byte
	// 对应select * 和  只有普通field情况
	if fetchCode == _allSource || fetchCode == _someSource {
		jsonMapper = func(hit *elastic.SearchHit) []byte {
			return hit.Source
		}
	} else if fetchCode == _id {
		// just _id
		jsonMapper = func(hit *elastic.SearchHit) []byte {
			bs := bytes.Buffer{}
			bs.WriteString("{\"_id\":\"")
			bs.WriteString(hit.Id)
			bs.WriteString("\"}")
			return bs.Bytes()
		}
	} else {
		// others
		jsonMapper = func(hit *elastic.SearchHit) []byte {
			row := GetSourceAndHighlight(hit)
			if Has(fetchCode, _index) {
				row["_index"] = hit.Index
			}
			if Has(fetchCode, _id) {
				row["_id"] = hit.Id
			}
			if Has(fetchCode, _score) {
				row["_score"] = hit.Score
			}
			bs, _ := json.Marshal(row)
			return bs
		}
	}
	return jsonMapper
}
func GetMapMapper(fetchCode int) func(hit *elastic.SearchHit) map[string]interface{} {
	var jsonMapper func(hit *elastic.SearchHit) map[string]interface{}
	// 对应select * 和  只有普通field情况
	if fetchCode == _allSource || fetchCode == _someSource {
		jsonMapper = func(hit *elastic.SearchHit) map[string]interface{} {
			row := make(map[string]interface{})
			if hit.Source != nil {
				json.Unmarshal(hit.Source, &row)
			}
			return row
		}
	} else if fetchCode == _id {
		// just _id
		jsonMapper = func(hit *elastic.SearchHit) map[string]interface{} {
			return map[string]interface{}{
				"_id": hit.Id,
			}
		}
	} else {
		// others
		jsonMapper = func(hit *elastic.SearchHit) map[string]interface{} {
			row := GetSourceAndHighlight(hit)
			if Has(fetchCode, _index) {
				row["_index"] = hit.Index
			}
			if Has(fetchCode, _id) {
				row["_id"] = hit.Id
			}
			if Has(fetchCode, _score) {
				row["_score"] = hit.Score
			}
			return row
		}
	}
	return jsonMapper
}

func GetSourceAndHighlight(hit *elastic.SearchHit) map[string]interface{} {
	row := make(map[string]interface{})
	if hit.Source != nil {
		json.Unmarshal(hit.Source, &row)
	}
	if hit.Highlight != nil {
		for f, vArray := range hit.Highlight {
			row[f] = vArray[0] // 只获一条
		}
	}
	return row
}
