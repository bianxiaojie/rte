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
	return ah.HandleNoneTargetAction(at, s, p).(Return)
}

func HandleOneTargetAction[Action OneTargetAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleOneTargetAction(at, s, id, p).(Return)
}

func HandleMutipleTargetsActionByIds[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleMutipleTargetsActionByIds(at, s, ids, p).(Return)
}

func HandleMutipleTargetsActionByIdPattern[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleMutipleTargetsActionByIdPattern(at, s, idPattern, p).(Return)
}

func HandleMutipleTargetsActionByType[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	tt := ref.ParseType[Target]()
	return ah.HandleMutipleTargetsActionByType(at, s, tt, p).(Return)
}

func HandleAllTargetsAction[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleAllTargetsAction(at, s, p).(Return)
}

func HandleNoneTargetStagedAction[Action NoneTargetStagedAction[Stage, Source, Param, Return], Stage ActionStage[Return], Source, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleNoneTargetStagedAction(at, s, p).(Return)
}

func HandleOneTargetStagedAction[Action OneTargetStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleOneTargetStagedAction(at, s, id, p).(Return)
}

func HandleMutipleTargetsStagedActionByIds[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleMutipleTargetsStagedActionByIds(at, s, ids, p).(Return)
}

func HandleMutipleTargetsStagedActionByIdPattern[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleMutipleTargetsStagedActionByIdPattern(at, s, idPattern, p).(Return)
}

func HandleMutipleTargetsStagedActionByType[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	tt := ref.ParseType[Target]()
	return ah.HandleMutipleTargetsStagedActionByType(at, s, tt, p).(Return)
}

func HandleAllTargetsStagedAction[Action MultipleTargetsStagedAction[Stage, Source, Target, Param, Return], Stage ActionStage[Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	at := ref.ParseType[Action]()
	return ah.HandleAllTargetsStagedAction(at, s, p).(Return)
}
