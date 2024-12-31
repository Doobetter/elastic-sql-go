package grammer

import (
	"github.com/olivere/elastic/v7"
	"strings"
)

type SortAdapter struct {
	Field string
	Ascending bool
	OrderLogic int; // asc 1 , desc -1, used for mem-sort
	Mode string
	NestedFilter Expression
	Script *ScriptAdapter
	ScriptSortType  string
}

func (a *SortAdapter) ToFieldSortBuilder() elastic.Sorter {
	var sort elastic.Sorter
	if a.Field!=""{
		if "_score" == a.Field{
			sort = elastic.NewScoreSort().Order(a.Ascending)

		}else {
			fsort:=elastic.NewFieldSort(a.Field).Order(a.Ascending)
			sort = fsort
			if a.Mode!=""{
				fsort.SortMode(a.Mode)
			}

			if strings.Contains(a.Field,"."){
				path := a.Field[0:strings.Index(a.Field,".")]
				nsort:=elastic.NewNestedSort(path)
				fsort.Nested(nsort)
				if a.NestedFilter!=nil{
					nsort.Filter(AdaptToQueryBuilder(a.NestedFilter,nil))
				}
			}

		}
	}else if a.Script!=nil{
		// todo
	}


	return sort
}