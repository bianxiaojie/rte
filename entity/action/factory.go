package action

import (
	"reflect"

	"github.com/bianxiaojie/rte/utils/ref"
)

type ActionFactory interface {
	GetAction(t reflect.Type) any
}

type defaultActionFactory struct {
	isSingleton        bool
	singletonActionMap map[reflect.Type]any
}

func MakeDefaultActionFactory() ActionFactory {
	return MakeDefaultActionFactoryWithParams(true)
}

func MakeDefaultActionFactoryWithParams(isSingleton bool) ActionFactory {
	af := &defaultActionFactory{}
	af.isSingleton = isSingleton
	if isSingleton {
		af.singletonActionMap = make(map[reflect.Type]any)
	}
	return af
}

func (af *defaultActionFactory) GetAction(t reflect.Type) any {
	if !af.isSingleton {
		return ref.New(t)
	}
	if _, ok := af.singletonActionMap[t]; !ok {
		af.singletonActionMap[t] = ref.New(t)
	}
	return af.singletonActionMap[t]
}

func GetNoneTargetAction[Action NoneTargetAction[Source, Param, Return], Source, Param, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}

func GetOneTargetAction[Action OneTargetAction[Source, Param, Target, Return], Source, Param, Target, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}

func GetMultipleTargetsAction[Action MultipleTargetsAction[Source, Param, Target, Return], Source, Param, Target, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}

func GetNoneTargetStagedAction[Action NoneTargetStagedAction[Stage, Source, Param, Return], Stage ActionStage[Return], Source, Param, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}

func GetOneTargetStagedAction[Action OneTargetStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}

func GetMultipleTargetsStagedAction[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](af ActionFactory) Action {
	return af.GetAction(ref.ParseType[Action]()).(Action)
}
