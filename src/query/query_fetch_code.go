package query

const (
	_index      = 1 << 0 // 1 获取_index
	_id         = 1 << 1 // 2 获取_id
	_score      = 1 << 2 // 4 获取 _score
	_allSource  = 1 << 3 // 8 获取所有的field 只获取全部field，而不获取其他,如_id
	_someSource = 1 << 4 // 16 获取部分field 只获取部分field，而不获取其他,如_id
	_hilight    = 1 << 5 // 32 获取高亮
	_script     = 1 << 6 // 64 获取_script_field
)

func Not(fetchCode int, code int) bool {
	return (fetchCode & code) == 0
}

func Has(fetchCode int, code int) bool {
	return (fetchCode & code) == code
}
