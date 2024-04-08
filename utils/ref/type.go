package ref

import "reflect"

func ParseType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}

func New(t reflect.Type) any {
	if t.Kind() == reflect.Pointer {
		return reflect.New(t.Elem()).Interface()
	}
	return reflect.New(t).Elem().Interface()
}
