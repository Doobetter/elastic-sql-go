/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package grammer

import "github.com/olivere/elastic/v7"

type HighlightAdapter struct {
	FieldAndSchema    map[string]string
	FieldSchema       []string // as
	Tag               string   // 与filed对应
	FragmentSize      int
	NumberOfFragments int
	NoMatchSize       int
}

func NewHighlightAdapter() *HighlightAdapter {
	// 设置默认值
	return &HighlightAdapter{
		FieldAndSchema:    make(map[string]string),
		FragmentSize:      100,
		NumberOfFragments: 1,
		NoMatchSize:       50,
	}
}

func (a *HighlightAdapter) ToHighlightBuilder() *elastic.Highlight {
	highlight := elastic.NewHighlight()
	for field, _ := range a.FieldAndSchema {
		hf := elastic.NewHighlighterField(field)

		highlight.Fields(hf)
	}
	//if a.Tag != "" {
	//	highlight.PreTags("<" + a.Tag + ">")
	//	highlight.PostTags("</" + a.Tag + ">")
	//	//highlight.PreTags( a.Tag )
	//	//highlight.PostTags(a.Tag)
	//}
	highlight.NumOfFragments(a.NumberOfFragments)
	highlight.FragmentSize(a.FragmentSize)
	highlight.NoMatchSize(a.NoMatchSize)
	return highlight
}
