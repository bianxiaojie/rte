package ref

import "reflect"

func ParseType[T any]() reflect.Type {
	return reflect.TypeOf((*T)(nil)).Elem()
}
