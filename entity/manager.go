package entity

import (
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/entity/behavior"
)

type EntityManager interface {
	EntityContainer
	behavior.BehaviorRepository
	action.ActionFactory
}

type defaultEntityManager struct {
	EntityContainer
	behavior.BehaviorRepository
	action.ActionFactory
}

func MakeDefaultEntityManager() EntityManager {
	em := &defaultEntityManager{}
	em.EntityContainer = MakeDefaultEntityContainer()
	em.BehaviorRepository = behavior.MakeDefaultBehaviorRepository()
	em.ActionFactory = action.MakeDefaultActionFactory()
	return em
}

func MakeDefaultEntityManagerWithParams(ec EntityContainer, br behavior.BehaviorRepository, af action.ActionFactory) EntityManager {
	em := &defaultEntityManager{
		EntityContainer:    ec,
		BehaviorRepository: br,
		ActionFactory:      af,
	}
	return em
}
