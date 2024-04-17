package action

type NoneTargetAction[Source, Param, Return any] interface {
	Action(Source, Param) Return
}

type OneTargetAction[Source, Target, Param, Return any] interface {
	Action(Source, Target, Param) Return
}

type MultipleTargetsAction[Source, Target, Param, Return any] interface {
	Action(Source, []Target, Param) Return
}

type ActionStage[Return any] interface {
	IsLastStage() bool
	GetReturnedValue() Return
}

type NoneTargetStagedAction[Stage ActionStage[Return], Source, Param, Return any] interface {
	MakeStage(Source, Param) Stage
	ActionStage(Source, Param, Stage) Stage
}

type OneTargetStagedAction[Stage ActionStage[Return], Source, Target, Param, Return any] interface {
	MakeStage(Source, Target, Param) Stage
	ActionStage(Source, Target, Param, Stage) Stage
}

type MultipleTargetsStagedAction[Stage ActionStage[Return], Source, Target, Param, Return any] interface {
	MakeStage(Source, []Target, Param) Stage
	ActionStage(Source, []Target, Param, Stage) Stage
}
