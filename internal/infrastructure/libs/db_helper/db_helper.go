package dbhelper

import (
	"fmt"
	"reflect"
)

func GetReflectName[T any](entity *T) string {
	return reflect.TypeOf(entity).Elem().Name()
}

func NestedTableName(tables ...string) string {
	templateS := tables[0]

	for index := range tables[1:] {
		templateS = fmt.Sprintf("%s.%s", templateS, tables[index+1])
	}

	return templateS
}
