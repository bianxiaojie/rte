package action

import (
	"reflect"

	"github.com/bianxiaojie/rte/utils/ref"
)

type ActionHandler interface {
	// non-staged version completes in one stage
	HandleNoneTargetAction(reflect.Type, any, any) any
	HandleOneTargetAction(reflect.Type, any, string, any) any
	HandleMutipleTargetsActionByIds(reflect.Type, any, []string, any) any
	HandleMutipleTargetsActionByIdPattern(reflect.Type, any, string, any) any
	HandleMutipleTargetsActionByType(reflect.Type, any, reflect.Type, any) any
	HandleAllTargetsAction(reflect.Type, any, any) any
	// staged version takes multiple stages
	HandleNoneTargetStagedAction(reflect.Type, any, any) any
	HandleOneTargetStagedAction(reflect.Type, any, string, any) any
	HandleMutipleTargetsStagedActionByIds(reflect.Type, any, []string, any) any
	HandleMutipleTargetsStagedActionByIdPattern(reflect.Type, any, string, any) any
	HandleMutipleTargetsStagedActionByType(reflect.Type, any, reflect.Type, any) any
	HandleAllTargetsStagedAction(reflect.Type, any, any) any
}

func HandleNoneTargetAction[Action NoneTargetAction[Source, Param, Return], Source, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleNoneTargetAction(at, s, p))
}

func HandleOneTargetAction[Action OneTargetAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleOneTargetAction(at, s, id, p))
}

func HandleMutipleTargetsActionByIds[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleMutipleTargetsActionByIds(at, s, ids, p))
}

func HandleMutipleTargetsActionByIdPattern[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleMutipleTargetsActionByIdPattern(at, s, idPattern, p))
}

func HandleMutipleTargetsActionByType[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	tt := ref.ParseType[Target]()
	return ref.Cast[Return](ah.HandleMutipleTargetsActionByType(at, s, tt, p))
}

func HandleAllTargetsAction[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleAllTargetsAction(at, s, p))
}

func HandleNoneTargetStagedAction[Action NoneTargetStagedAction[Stage, Source, Param, Return], Stage ActionStage[Return], Source, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleNoneTargetStagedAction(at, s, p))
}

func HandleOneTargetStagedAction[Action OneTargetStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleOneTargetStagedAction(at, s, id, p))
}

func HandleMutipleTargetsStagedActionByIds[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleMutipleTargetsStagedActionByIds(at, s, ids, p))
}

func HandleMutipleTargetsStagedActionByIdPattern[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleMutipleTargetsStagedActionByIdPattern(at, s, idPattern, p))
}

func HandleMutipleTargetsStagedActionByType[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	tt := ref.ParseType[Target]()
	return ref.Cast[Return](ah.HandleMutipleTargetsStagedActionByType(at, s, tt, p))
}

func HandleAllTargetsStagedAction[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ref.Cast[Return](ah.HandleAllTargetsStagedAction(at, s, p))
}
