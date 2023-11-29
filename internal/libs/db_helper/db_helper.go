package dbhelper

import (
	"reflect"
)

func GetEntityTableName[T any](entity *T) string {
	return reflect.TypeOf(entity).Elem().Name()
}
