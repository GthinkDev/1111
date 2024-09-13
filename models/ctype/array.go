package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (a *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) != "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

func (a Array) Value() (driver.Value, error) {
	// 将数字切片转换为字符串切片
	return strings.Join(a, "\n"), nil
}
