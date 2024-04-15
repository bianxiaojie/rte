package behavior

import (
	"reflect"
	"slices"
)

type BehaviorRepository interface {
	AddBehavior(Behavior)
	AddBehaviorByType(reflect.Type)
	RemoveBehavior(Behavior)
	RemoveBehaviorByType(reflect.Type)
	GetSortedBehaviors() [][]Behavior
}

type defaultBehaviorRepository struct {
	behaviorMap           map[reflect.Type][]Behavior
	behaviorParser        BehaviorParser
	cachedSortedBehaviors [][]Behavior
}

func MakeDefaultBehaviorRepository() BehaviorRepository {
	return MakeDefaultBehaviorRepositoryWithParams(MakeDefaultBehaviorParser())
}

func MakeDefaultBehaviorRepositoryWithParams(bp BehaviorParser) BehaviorRepository {
	br := &defaultBehaviorRepository{}
	br.behaviorMap = make(map[reflect.Type][]Behavior)
	br.behaviorParser = bp
	return br
}

func (br *defaultBehaviorRepository) AddBehavior(b Behavior) {
	rt := b.ReceiverType()
	behaviors, ok := br.behaviorMap[rt]
	if !ok {
		behaviors = make([]Behavior, 0)
	} else if slices.ContainsFunc(behaviors, b.Equal) {
		return
	}

	br.behaviorMap[rt] = append(behaviors, b)
	br.cachedSortedBehaviors = nil
}

func (br *defaultBehaviorRepository) AddBehaviorByType(t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		if behavior, ok := br.behaviorParser.ParseBehavior(t.Method(i)); ok {
			br.AddBehavior(behavior)
		}
	}
}

func (br *defaultBehaviorRepository) RemoveBehavior(b Behavior) {
	rt := b.ReceiverType()
	behaviors, ok := br.behaviorMap[rt]
	if !ok || !slices.ContainsFunc(behaviors, b.Equal) {
		return
	}

	br.behaviorMap[rt] = slices.DeleteFunc(behaviors, b.Equal)
	if len(br.behaviorMap[rt]) == 0 {
		delete(br.behaviorMap, rt)
	}
	br.cachedSortedBehaviors = nil
}

func (br *defaultBehaviorRepository) RemoveBehaviorByType(t reflect.Type) {
	for i := 0; i < t.NumMethod(); i++ {
		if behavior, ok := br.behaviorParser.ParseBehavior(t.Method(i)); ok {
			br.RemoveBehavior(behavior)
		}
	}
}

func (br *defaultBehaviorRepository) GetSortedBehaviors() [][]Behavior {
	if br.cachedSortedBehaviors == nil {
		allBehaviors := make([]Behavior, 0)
		for _, behaviors := range br.behaviorMap {
			allBehaviors = append(allBehaviors, behaviors...)
		}
		br.cachedSortedBehaviors = SortBehaviorsByPriority(allBehaviors)
	}

	return br.cachedSortedBehaviors
}
