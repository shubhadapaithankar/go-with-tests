package main

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}
	return val
}
