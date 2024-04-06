package behavior

import (
	"reflect"

	"github.com/bianxiaojie/rte/utils/ref"
)

type Behavior interface {
	Name() string
	EntityType() reflect.Type
	Priority() int // a lower number means a higher priority
	Call(reciever any, args ...any) []any
}

type defaultBehavior struct {
	method     reflect.Method
	name       string
	entityType reflect.Type
	priority   int
}

func (b defaultBehavior) Name() string {
	return b.name
}

func (b defaultBehavior) Priority() int {
	return b.priority
}

func (b defaultBehavior) EntityType() reflect.Type {
	return b.entityType
}

func (b defaultBehavior) Call(reciever any, args ...any) []any {
	return ref.CallMethod(b.method, reciever, args...)
}
