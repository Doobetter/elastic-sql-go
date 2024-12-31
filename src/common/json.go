package common

import "encoding/json"

func JSONStr(obj interface{})string{
	bs, err:=json.Marshal(obj)
	if err!=nil{
		return ""
	}
	return string(bs)
}

func JSONPrettyStr(obj interface{})string{
	bs, err:=json.MarshalIndent(obj,"","  ")
	if err!=nil{
		return ""
	}
	return string(bs)
}
func JSONPrettyStrWithErr(obj interface{}, err error)string{
	if err!=nil{
		return ""
	}
	bs, err:=json.MarshalIndent(obj,"","  ")
	if err!=nil{
		return ""
	}
	return string(bs)
}
func JSONStrWith(obj interface{}, err error)string{
	if err!=nil{
		return ""
	}
	bs, err:=json.Marshal(obj)
	if err!=nil{
		return err.Error()
	}
	return string(bs)
}
