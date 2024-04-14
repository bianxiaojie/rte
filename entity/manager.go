package entity

import (
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/entity/behavior"
)

type EntityManager struct {
	EntityContainer
	behavior.BehaviorRepository
	action.ActionFactory
}

func MakeDefaultEntityManager() *EntityManager {
	em := &EntityManager{}
	em.EntityContainer = MakeDefaultEntityContainer()
	em.BehaviorRepository = behavior.MakeDefaultBehaviorRepository()
	em.ActionFactory = action.MakeDefaultActionFactory()
	return em
}

func MakeEntityManager(ec EntityContainer, br behavior.BehaviorRepository, af action.ActionFactory) *EntityManager {
	em := &EntityManager{
		EntityContainer:    ec,
		BehaviorRepository: br,
		ActionFactory:      af,
	}
	return em
}
