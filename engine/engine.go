package engine

import (
	"time"

	"github.com/bianxiaojie/rte/engine/stage"
	"github.com/bianxiaojie/rte/entity"
	"github.com/bianxiaojie/rte/scheduler"
	"github.com/bianxiaojie/rte/timer"
)

type Engine interface {
	Start()
	Pause()
	Resume()
	Stop()
	WaitStopped()
	EntityManager() entity.EntityManager
}

type defaultEngine struct {
	it timer.IncrementalTimer
	s  scheduler.Scheduler
	em entity.EntityManager
}

func MakeDefaultEngine(unit time.Duration) Engine {
	e := &defaultEngine{}
	e.it = timer.MakeDefaultIncrementalTimer(unit)
	e.s = scheduler.MakeDefaultScheduler()
	e.em = entity.MakeDefaultEntityManager()
	return e
}

func (e *defaultEngine) Start() {
	stageTimer := stage.MakeStageTimer(e.s, e.it)
	stageEntity := stage.MakeStageEntity(e.em, e.it)
	e.s.Start(stageTimer, stageEntity)
}

func (e *defaultEngine) Pause() {
	e.s.Pause()
}

func (e *defaultEngine) Resume() {
	e.s.Resume()
}

func (e *defaultEngine) Stop() {
	e.s.Stop()
}

func (e *defaultEngine) WaitStopped() {
	e.s.WaitStopped()
}

func (e *defaultEngine) EntityManager() entity.EntityManager {
	return e.em
}
