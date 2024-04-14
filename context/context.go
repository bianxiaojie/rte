package context

import (
	"github.com/bianxiaojie/rte/entity"
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/timer"
)

type Context interface {
	SetParam(string, any)
	GetParam(string, any)
	RemoveParam(string, any)
	EntityManager() entity.EntityManager
	ActionHandler() action.ActionHandler
	Timer() timer.Timer
}
