package action

type ActionHandler interface {
	// non-staged version completes in one stage
	HandleNoneTargetAction(any, any) any
	HandleOneTargetAction(any, string, any) any
	HandleMutipleTargetsActionByIds(any, []string, any) any
	HandleMutipleTargetsActionByIdPattern(any, string, any) any
	HandleAllTargetsAction(any, any) any
	// staged version takes multiple stages
	HandleNoneTargetStagedAction(any, any) any
	HandleOneTargetStagedAction(any, string, any) any
	HandleMutipleTargetsStagedActionByIds(any, []string, any) any
	HandleMutipleTargetsStagedActionByIdPattern(any, string, any) any
	HandleAllTargetsStagedAction(any, any) any
}

func HandleNoneTargetAction[Action NoneTargetAction[Source, Param, Return], Source, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	return ah.HandleNoneTargetAction(s, p).(Return)
}

func HandleOneTargetAction[Action OneTargetAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	return ah.HandleOneTargetAction(s, id, p).(Return)
}

func HandleMutipleTargetsActionByIds[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	return ah.HandleMutipleTargetsActionByIds(s, ids, p).(Return)
}

func HandleMutipleTargetsActionByIdPattern[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	return ah.HandleMutipleTargetsActionByIdPattern(s, idPattern, p).(Return)
}

func HandleAllTargetsAction[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	return ah.HandleAllTargetsAction(s, p).(Return)
}

func HandleNoneTargetStagedAction[Action NoneTargetAction[Source, Param, Return], Source, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	return ah.HandleNoneTargetStagedAction(s, p).(Return)
}

func HandleOneTargetStagedAction[Action OneTargetAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, id string, p Param) Return {
	return ah.HandleOneTargetStagedAction(s, id, p).(Return)
}

func HandleMutipleTargetsStagedActionByIds[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, ids []string, p Param) Return {
	return ah.HandleMutipleTargetsStagedActionByIds(s, ids, p).(Return)
}

func HandleMutipleTargetsStagedActionByIdPattern[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, idPattern string, p Param) Return {
	return ah.HandleMutipleTargetsStagedActionByIdPattern(s, idPattern, p).(Return)
}

func HandleAllTargetsStagedAction[Action MultipleTargetsAction[Source, Target, Param, Return], Source, Target, Param, Return any](ah ActionHandler, s Source, p Param) Return {
	return ah.HandleAllTargetsStagedAction(s, p).(Return)
}
