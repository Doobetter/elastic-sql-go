/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package basic

type Statement interface {
	Init(ctx *ExeElasticSQLCtx) error

	Execute(ctx *ExeElasticSQLCtx) error

	Clean(ctx *ExeElasticSQLCtx)
	//GetExportFileName 如果是可导出的Statement返回导出文件名，子类可以覆盖实现
	GetExportFileName() string
	DefaultNameSuffix() string
	GetName() string
	SetName(name string)
	GenPostProcessCode() int
}

//
//func (s *Statement) init(ctx *ExeElasticSQLCtx) error { return nil }
//
//func (s *Statement) execute(ctx *ExeElasticSQLCtx) error { return nil }
//
//func (s *Statement) clean(ctx *ExeElasticSQLCtx) {}
//
//func (s *Statement) print() {}
//
////GetExportFileName 如果是可导出的Statement返回导出文件名，子类可以覆盖实现
//func (s *Statement) GetExportFileName() string {
//	return ""
//}
