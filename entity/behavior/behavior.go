package behavior

import (
	"reflect"
	"slices"

	"github.com/bianxiaojie/rte/utils/ref"
)

type Behavior interface {
	Name() string
	ReceiverType() reflect.Type
	Priority() int // a lower number means a higher priority
	Call(reciever any, args ...any) []any
	Equal(Behavior) bool
}

func SortBehaviorsByPriority(behaviors []Behavior) [][]Behavior {
	sortedBehaviors := make([][]Behavior, 0)
	if len(behaviors) == 0 {
		return sortedBehaviors
	}

	slices.SortStableFunc(behaviors, func(a, b Behavior) int {
		return a.Priority() - b.Priority()
	})

	i := 0
	sortedBehaviors = append(sortedBehaviors, []Behavior{behaviors[0]})
	for j := 1; j < len(behaviors); j++ {
		if behaviors[j].Priority() != behaviors[j-1].Priority() {
			i++
			sortedBehaviors = append(sortedBehaviors, []Behavior{})
		}
		sortedBehaviors[i] = append(sortedBehaviors[i], behaviors[j])
	}

	return sortedBehaviors
}

type defaultBehavior struct {
	method       reflect.Method
	name         string
	recieverType reflect.Type
	priority     int
}

func (b *defaultBehavior) Name() string {
	return b.name
}

func (b *defaultBehavior) Priority() int {
	return b.priority
}

func (b *defaultBehavior) ReceiverType() reflect.Type {
	return b.recieverType
}

func (b *defaultBehavior) Call(reciever any, args ...any) []any {
	return ref.CallMethod(b.method, reciever, args...)
}

func (b *defaultBehavior) Equal(behavior Behavior) bool {
	return b.Name() == behavior.Name() && b.Priority() == behavior.Priority()
}
