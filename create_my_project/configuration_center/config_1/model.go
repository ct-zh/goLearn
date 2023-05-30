package config_1

type KeyValStruct struct {
	Id        int64      `json:"id" gorm:"column:id"`
	Key       string     `json:"key" gorm:"column:key"`
	Value     string     `json:"value" gorm:"column:value"`
	Status    int8       `json:"status" gorm:"status"`
	ValueType KeyValType `json:"value_type" gorm:"value_type"`
}

type KeyValType int8

var KeyValTypes = struct {
	Text KeyValType
	JSON KeyValType
}{
	Text: 1,
	JSON: 2,
}

func (KeyValStruct) getKeyValTable() string {
	return "key_val"
}
