package stage

import (
	"fmt"
	"reflect"

	"github.com/bianxiaojie/rte/context"
	"github.com/bianxiaojie/rte/entity"
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/entity/behavior"
	"github.com/bianxiaojie/rte/scheduler"
	"github.com/bianxiaojie/rte/timer"
	"github.com/bianxiaojie/rte/utils/ref"
)

type defaultActionHandler struct {
	linkId string
	em     entity.EntityManager
	sl     scheduler.SchedulerLinker
	se     *StageEntity
}

func makeDefaultActionHandler(linkId string, em entity.EntityManager, sl scheduler.SchedulerLinker, se *StageEntity) action.ActionHandler {
	ah := &defaultActionHandler{}
	ah.linkId = linkId
	ah.em = em
	ah.sl = sl
	ah.se = se
	return ah
}

func (ah *defaultActionHandler) HandleNoneTargetAction(actionType reflect.Type, source any, param any) any {
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, source, param)
}

func (ah *defaultActionHandler) HandleOneTargetAction(actionType reflect.Type, source any, id string, param any) any {
	target, _ := ah.em.GetEntityById(id)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, action, source, target, param)
}

func (ah *defaultActionHandler) HandleMutipleTargetsActionByIds(actionType reflect.Type, source any, ids []string, param any) any {
	targets := ah.em.GetEntitiesByIds(ids)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, action, source, targets, param)
}

func (ah *defaultActionHandler) HandleMutipleTargetsActionByIdPattern(actionType reflect.Type, source any, idPattern string, param any) any {
	targets := ah.em.GetEntitiesByIdPattern(idPattern)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, action, source, targets, param)
}

func (ah *defaultActionHandler) HandleMutipleTargetsActionByType(actionType reflect.Type, source any, targetType reflect.Type, param any) any {
	targets := ah.em.GetEntitiesByType(targetType)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, action, source, targets, param)
}

func (ah *defaultActionHandler) HandleAllTargetsAction(actionType reflect.Type, source any, param any) any {
	targets := ah.em.GetEntities()
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("Action")
	return ref.CallFunc(method, action, source, targets, param)
}

func (ah *defaultActionHandler) HandleNoneTargetStagedAction(actionType reflect.Type, source any, param any) any {
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          none,
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

func (ah *defaultActionHandler) HandleOneTargetStagedAction(actionType reflect.Type, source any, id string, param any) any {
	target, _ := ah.em.GetEntityById(id)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          one,
		targets:     []entity.Entity{target},
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

func (ah *defaultActionHandler) HandleMutipleTargetsStagedActionByIds(actionType reflect.Type, source any, ids []string, param any) any {
	targets := ah.em.GetEntitiesByIds(ids)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          multiple,
		targets:     targets,
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

func (ah *defaultActionHandler) HandleMutipleTargetsStagedActionByIdPattern(actionType reflect.Type, source any, idPattern string, param any) any {
	targets := ah.em.GetEntitiesByIdPattern(idPattern)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          multiple,
		targets:     targets,
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

func (ah *defaultActionHandler) HandleMutipleTargetsStagedActionByType(actionType reflect.Type, source any, targetType reflect.Type, param any) any {
	targets := ah.em.GetEntitiesByType(targetType)
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          multiple,
		targets:     targets,
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

func (ah *defaultActionHandler) HandleAllTargetsStagedAction(actionType reflect.Type, source any, param any) any {
	targets := ah.em.GetEntities()
	action := ah.em.GetAction(actionType)
	method := reflect.ValueOf(action).MethodByName("MakeStage")
	actionStage := ref.CallFunc(method, param)[0]
	sai := &stagedActionInfo{
		action:      action,
		source:      source,
		at:          multiple,
		targets:     targets,
		param:       param,
		actionStage: actionStage,
	}
	ah.se.setStageActionInfo(ah.linkId, sai)
	return ah.sl.LinkAndWait(ah.linkId)
}

type actionType int

const (
	none actionType = iota
	one
	multiple
)

type stagedActionInfo struct {
	action      any
	source      any
	at          actionType
	targets     []entity.Entity
	param       any
	actionStage any
}

type StageEntity struct {
	gid                       int64
	em                        entity.EntityManager
	t                         timer.Timer
	sortedBehaviors           [][]behavior.Behavior
	entities                  []entity.Entity
	indices                   [3]int
	linkId2StageActionInfoMap map[string]*stagedActionInfo
}

func MakeStageEntity(em entity.EntityManager, t timer.Timer) *StageEntity {
	se := &StageEntity{}
	se.em = em
	se.t = t
	se.linkId2StageActionInfoMap = make(map[string]*stagedActionInfo)
	return se
}

func (se *StageEntity) Run(sl scheduler.SchedulerLinker) {
	se.gid = sl.GoroutineId()

	// init behaviors process
	if se.sortedBehaviors == nil {
		se.sortedBehaviors = se.em.GetSortedBehaviors()
	}

	// call all behaviors
	for se.indices[0] < len(se.sortedBehaviors) {
		i := se.indices[0]
		for se.indices[1] < len(se.sortedBehaviors[i]) {
			j := se.indices[1]
			behavior := se.sortedBehaviors[i][j]
			if se.entities == nil {
				se.entities = se.em.GetEntitiesByType(behavior.ReceiverType())
			}
			for se.indices[2] < len(se.entities) {
				entity := se.entities[se.indices[2]]
				if !se.processBehavior(sl, behavior, entity) {
					return
				}
				se.indices[2]++
			}
			se.entities = nil
			se.indices[1]++
			se.indices[2] = 0
		}
		se.indices[0]++
		se.indices[1] = 0
	}
	se.sortedBehaviors = nil
	se.indices = [3]int{0, 0, 0}
}

func (se *StageEntity) processBehavior(sl scheduler.SchedulerLinker, b behavior.Behavior, e entity.Entity) bool {
	linkId := fmt.Sprintf("%s-%s-%s", b.ReceiverType().Name(), b.Name(), e.Id())
	if sai, ok := se.linkId2StageActionInfoMap[linkId]; ok {
		method := reflect.ValueOf(sai.action).MethodByName("ActionStage")
		switch sai.at {
		case none:
			sai.actionStage = ref.CallFunc(method, sai.source, sai.actionStage)[0]
		case one:
			sai.actionStage = ref.CallFunc(method, sai.source, sai.targets[0], sai.actionStage)[0]
		case multiple:
			sai.actionStage = ref.CallFunc(method, sai.source, sai.targets, sai.actionStage)[0]
		}
		method = reflect.ValueOf(sai.actionStage).MethodByName("IsLastStage")
		if ref.CallFunc(method)[0].(bool) {
			method = reflect.ValueOf(sai.actionStage).MethodByName("GetReturnedValue")
			sl.SetResultAndWait(linkId, ref.CallFunc(method)[0])
		}
	} else {
		ah := makeDefaultActionHandler(linkId, se.em, sl, se)
		ctx := context.MakeDefaultContext(se.em, ah, se.t)
		b.Call(e, ctx)
		if se.gid != sl.GoroutineId() {
			delete(se.linkId2StageActionInfoMap, linkId)
			sl.Unlink(linkId)

			return false
		}
	}

	return true
}

func (se *StageEntity) setStageActionInfo(linkId string, sai *stagedActionInfo) {
	if _, ok := se.linkId2StageActionInfoMap[linkId]; ok {
		panic(fmt.Sprintf("one entity cannot action in mutiple same behaviors at the same time: %s", linkId))
	}
	se.linkId2StageActionInfoMap[linkId] = sai
}
