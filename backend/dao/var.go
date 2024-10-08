/**
 ******************************************************************************
 * @file           : var.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/8/24
 ******************************************************************************
 */

package dao

import (
	_interface "backend/dao/interface"
	"backend/dao/model"
	"reflect"
)

func getStructName(i interface{}) string {
	return reflect.TypeOf(i).Name()
}

func getStructInstance(t reflect.Type) interface{} {
	return reflect.New(t).Elem().Interface()
}

var types = []reflect.Type{
	reflect.TypeOf(model.Danmu{}),
	reflect.TypeOf(model.User{}),
	reflect.TypeOf(model.Category{}),
	reflect.TypeOf(model.Article{}),
	reflect.TypeOf(model.Comment{}),
	reflect.TypeOf(model.History{}),
	reflect.TypeOf(model.Live{}),
	reflect.TypeOf(model.Video{}),
	reflect.TypeOf(model.Tag{}),
}

var mps = map[string][]interface{}{
	getStructName(model.User{}): {func(user _interface.UserDaoInterface) {}},
}
