package behavior

import (
	"reflect"
	"strconv"

	"github.com/bianxiaojie/rte/utils/str"
)

type BehaviorParser interface {
	ParseBehavior(method reflect.Method) (Behavior, bool)
}

type defaultBehaviorParser struct{}

func MakeDefaultBehaviorParser() BehaviorParser {
	return defaultBehaviorParser{}
}

func (bp defaultBehaviorParser) ParseBehavior(method reflect.Method) (Behavior, bool) {
	// behavior method name must be like F_1
	groups, ok := str.MatchAndFindGroups(`^([a-zA-Z0-9]+)_(\d+)$`, method.Name)
	if !ok {
		return nil, false
	}

	name, priority := groups[1], groups[2]
	priorityInt, _ := strconv.Atoi(priority)

	b := defaultBehavior{
		method:     method,
		name:       name,
		entityType: method.Type.In(0),
		priority:   priorityInt,
	}
	return b, true
}
