package entity

import (
	"reflect"
	"regexp"

	"github.com/bianxiaojie/rte/utils/sli"
)

type EntityContainer interface {
	AddEntity(Entity)
	RemoveEntityById(string)
	RemoveEntitiesByIdPattern(string)
	RemoveEntitiesByType(reflect.Type)
	RemoveEntities()
	GetEntityById(string) (Entity, bool)
	GetEntitiesByIdPattern(string) []Entity
	GetEntitiesByType(reflect.Type) []Entity
	GetEntities() []Entity
}

type defaultEntityContainer struct {
	id2EntityMap map[string]Entity
	type2IdMap   map[reflect.Type][]string
}

func MakeDefaultEntityContainer() EntityContainer {
	ec := &defaultEntityContainer{}
	ec.id2EntityMap = make(map[string]Entity)
	ec.type2IdMap = make(map[reflect.Type][]string)
	return ec
}

func (ec *defaultEntityContainer) AddEntity(e Entity) {
	id := e.Id()
	existedEntity, ok := ec.id2EntityMap[id]
	if ok && existedEntity != e {
		panic("cannot override different entity with same id")
	}
	if ok && existedEntity == e {
		return
	}

	ec.id2EntityMap[id] = e
	et := reflect.TypeOf(e)
	ec.type2IdMap[et] = append(ec.type2IdMap[et], id)
}

func (ec *defaultEntityContainer) RemoveEntityById(id string) {
	e, ok := ec.id2EntityMap[id]
	if !ok {
		return
	}

	et := reflect.TypeOf(e)
	ec.type2IdMap[et] = sli.Delete(ec.type2IdMap[et], id)
	if len(ec.type2IdMap[et]) == 0 {
		delete(ec.type2IdMap, et)
	}
	delete(ec.id2EntityMap, id)
}

func (ec *defaultEntityContainer) RemoveEntitiesByIdPattern(pattern string) {
	candidateIds := make([]string, 0)
	for id := range ec.id2EntityMap {
		if ok, _ := regexp.MatchString(pattern, id); ok {
			candidateIds = append(candidateIds, id)
		}
	}

	for _, id := range candidateIds {
		ec.RemoveEntityById(id)
	}
}

func (ec *defaultEntityContainer) RemoveEntitiesByType(et reflect.Type) {
	candidateIds, ok := ec.type2IdMap[et]
	if !ok {
		return
	}

	for _, id := range candidateIds {
		ec.RemoveEntityById(id)
	}
}

func (ec *defaultEntityContainer) RemoveEntities() {
	ec.id2EntityMap = make(map[string]Entity)
	ec.type2IdMap = make(map[reflect.Type][]string)
}

func (ec *defaultEntityContainer) GetEntityById(id string) (Entity, bool) {
	entity, ok := ec.id2EntityMap[id]
	return entity, ok
}

func (ec *defaultEntityContainer) GetEntitiesByIdPattern(pattern string) []Entity {
	entities := make([]Entity, 0)
	for id, entity := range ec.id2EntityMap {
		if ok, _ := regexp.MatchString(pattern, id); ok {
			entities = append(entities, entity)
		}
	}
	return entities
}

func (ec *defaultEntityContainer) GetEntitiesByType(et reflect.Type) []Entity {
	ids, ok := ec.type2IdMap[et]
	if !ok {
		return []Entity{}
	}

	entites := make([]Entity, len(ids))
	for i, id := range ids {
		entites[i] = ec.id2EntityMap[id]
	}
	return entites
}

func (ec *defaultEntityContainer) GetEntities() []Entity {
	entities := make([]Entity, 0)
	for _, entity := range ec.id2EntityMap {
		entities = append(entities, entity)
	}
	return entities
}
