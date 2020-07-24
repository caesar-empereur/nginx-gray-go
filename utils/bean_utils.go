package utils

import (
	"fmt"
	"reflect"
)

func CopyFields(src interface{}, tar interface{}) {
	srcType := reflect.TypeOf(src)
	srcValue := reflect.ValueOf(src)
	tarType := reflect.TypeOf(tar)
	tarValue := reflect.ValueOf(tar)

	// srcType.Kind() 返回值包括 struct，Ptr
	if srcType.Kind() != reflect.Ptr {
		panic("src must be a struct pointer")
	}
	tarFields := make([]string, 0)
	for i := 0; i < tarValue.Elem().NumField(); i++ {
		tarFields = append(tarFields, tarType.Elem().Field(i).Name)
	}

	if len(tarFields) == 0 {
		panic("no fields to copy")
	}

	// srcValue.Elem() 获取指针指向的元素
	for i := 0; i < len(tarFields); i++ {
		name := tarFields[i]
		srcFieldValue := srcValue.Elem().FieldByName(name)
		tarFieldValue := tarValue.Elem().FieldByName(name)

		if tarFieldValue.IsValid() && srcFieldValue.Kind() == tarFieldValue.Kind() {
			tarFieldValue.Set(srcFieldValue)
		} else {
			fmt.Printf("no such field or different kind, fieldName: %s\n", name)
		}
	}
	return
}
