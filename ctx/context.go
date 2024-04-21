package ctx

import (
	"github.com/bianxiaojie/rte/entity"
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/timer"
)

type Context interface {
	SetParam(string, any)
	GetParam(string) (any, bool)
	RemoveParam(string) bool
	EntityManager() entity.EntityManager
	ActionHandler() action.ActionHandler
	Timer() timer.Timer
}

type defaultContext struct {
	paramId2paramMap map[string]any
	em               entity.EntityManager
	ah               action.ActionHandler
	t                timer.Timer
}

func MakeDefaultContext(em entity.EntityManager, ah action.ActionHandler, t timer.Timer) Context {
	c := &defaultContext{}
	c.paramId2paramMap = make(map[string]any)
	c.em = em
	c.ah = ah
	c.t = t
	return c
}

func (c *defaultContext) SetParam(paramId string, param any) {
	c.paramId2paramMap[paramId] = param
}

func (c *defaultContext) GetParam(paramId string) (any, bool) {
	param, ok := c.paramId2paramMap[paramId]
	return param, ok
}

func (c *defaultContext) RemoveParam(paramId string) bool {
	_, ok := c.paramId2paramMap[paramId]
	if ok {
		delete(c.paramId2paramMap, paramId)
	}
	return ok
}

func (c *defaultContext) EntityManager() entity.EntityManager {
	return c.em
}

func (c *defaultContext) ActionHandler() action.ActionHandler {
	return c.ah
}

func (c *defaultContext) Timer() timer.Timer {
	return c.t
}
