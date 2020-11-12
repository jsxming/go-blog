package util

type TableData struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

func NewTableData(total int64, list interface{}) TableData {
	return TableData{
		Total: total,
		List:  list,
	}
}
