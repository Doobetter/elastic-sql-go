/*
 *  Copyright 2020-present Doobetter. All rights reserved.
 *  Use of this source code is governed by a MIT-license.
 *
 */

package basic

import "strconv"

type ProcessUnit struct {
	Name string //  Name as label，处理单元名字

	postProcessesCode int // 后续处理编码 默认1

	Processes map[int]Void // 该处理单元的所有后续处理过程
}

func NewProcessUnit() *ProcessUnit {
	p := new(ProcessUnit)
	p.Processes = make(map[int]Void)
	return p
}

func (p *ProcessUnit) GetName() string {
	return p.Name
}
func (p *ProcessUnit) SetName(name string) {
	p.Name = name
}

// DefaultNameSuffix 默认名的缀
func (p *ProcessUnit) DefaultNameSuffix() string {
	return ""
}

// GenUniqName 生成处理单元的名字
// 如果没有指定名字，系统生成唯一的名字
func (p *ProcessUnit) GenUniqName(mapName string, seq int) string {
	if mapName == "" {
		p.Name = strconv.Itoa(seq) + p.DefaultNameSuffix()
	} else {
		p.Name = mapName
	}
	return p.Name
}

func (p *ProcessUnit) Add(code int) {
	p.Processes[code] = VoidValue
}
func (p *ProcessUnit) GenPostProcessCode() int {
	// default is 1
	code := 1
	if len(p.Processes) > 0 {
		code = 0
		for k, _ := range p.Processes {
			code = code | k
		}

	}
	p.postProcessesCode = code

	return p.postProcessesCode
}

// IsViaNON 是否后续没有任何操作使用该过程结果
func (p *ProcessUnit) IsViaNON() bool {
	return p.postProcessesCode == PostProcessEnumVIANON
}

// IsViaExport 该过程的结果是否要用于 export 输出结果
func (p *ProcessUnit) IsViaExport() bool {
	return (PostProcessEnumVIAEXPORT & p.postProcessesCode) == PostProcessEnumVIAEXPORT
}

// IsJustViaExport 是否只需要export
func (p *ProcessUnit) IsJustViaExport() bool {
	return p.postProcessesCode == PostProcessEnumVIAEXPORT
}

// IsViaJOIN 该过程的结果是否要用于 join statement
func (p *ProcessUnit) IsViaJOIN() bool {
	return (PostProcessEnumVIAJOIN & p.postProcessesCode) == PostProcessEnumVIAJOIN
}

// IsViaSpark 该过程的结果是否要用于 spark statement
func (p *ProcessUnit) IsViaSpark() bool {
	return (PostProcessEnumVIASPARK & p.postProcessesCode) == PostProcessEnumVIASPARK
}

// IsViaHive 该过程的结果是否要用于 hive sql
func (p *ProcessUnit) IsViaHive() bool {
	return (PostProcessEnumVIAHIVE & p.postProcessesCode) == PostProcessEnumVIAHIVE
}

// IsViaByAnd 是否都要经过vias后续处理
func (p *ProcessUnit) IsViaByAnd(vias ...int) bool {
	code := 0
	for _, via := range vias {
		code = code | via
	}
	return code == p.postProcessesCode
}

// IsViaByOr 只要经过vias中的一个就返回true
func (p *ProcessUnit) IsViaByOr(vias ...int) bool {
	code := 0
	for _, via := range vias {
		code = code | via
	}
	return (code & p.postProcessesCode) != 0
}

const (
	PostProcessEnumVIANON = 1 // 默认只查询

	PostProcessEnumVIAEXPORT = 2 // 需要输出

	PostProcessEnumVIAJOIN = 4 // 需要ES_JOIN

	PostProcessEnumVIASPARK = 8 // 需要经过SPARK处理

	PostProcessEnumVIAHIVE = 16 // 需要经过HIVE处理
)

type Void struct {
}

var VoidValue Void
