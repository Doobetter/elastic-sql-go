/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package basic

import "github.com/olivere/elastic/v7"

type ResultSet struct {
	Name           string                   `json:"name,omitempty"`
	Count          int64                    `json:"count,omitempty"`
	Total          int64                    `json:"total,omitempty"`
	FetchSize      int64                    `json:"fetch_size,omitempty"`
	ErrMsg         string                   `json:"err_msg,omitempty"`
	WarnMsg        string                   `json:"warn_msg,omitempty"`
	DataStr        string                   `json:"data_str,omitempty"` //存储用于输出csv/json的数据
	ScrollId       string                   `json:"scroll_id,omitempty"`
	Schemas        []string                 `json:"schemas,omitempty"`
	Data           []map[string]interface{} `json:"data,omitempty"` // 存储数据
	Headers        []string                 `json:"headers,omitempty"`
	SearchResponse *elastic.SearchResult    `json:"-"`
}

func NewResultSet(name string) *ResultSet {
	return &ResultSet{
		Name:     name,
		Count:    0,
		Total:    0,
		ErrMsg:   "",
		DataStr:  "",
		ScrollId: "",
		Schemas:  nil,
		Data:     nil,
	}
}

type EmptyResultSet struct {
	ResultSet
}

func NewEmptyResultSet(name string) *EmptyResultSet {
	rs := new(EmptyResultSet)
	rs.Name = name
	return rs
}
