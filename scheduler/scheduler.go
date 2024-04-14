package scheduler

import (
	"fmt"
	"sync"
)

type SchedulerLinker interface {
	// id of current goroutine
	GoroutineId() int64
	// fork a linked goroutine or signal existed one to continue scheduling stages.
	// Then current goroutine blocks itself and waits for the linked goroutine to set result and signal it.
	// After being signalled, it can Unlink or LinkAndWait again.
	// Note that once a goroutine LinkAndWait, it can
	// input string is the linkId used to link two goroutines.
	LinkAndWait(string) any
	// set the result and signal the linked goroutine, then wait for the linked goroutine to signal it.
	// input string is the linkId used to link two goroutines.
	SetResultAndWait(string, any)
	// unlink current goroutine with linked goroutine, then signal linked goroutine and terminate itself.
	Unlink(string)
	// here is a timeline for LinkAndWait and SetResultAndWait.
	// A: running -> LinkAndWait() -> blocking                       -> Unlink()/LinkAndWait()
	// B:                          -> running  -> SetResultAndWait() -> blocking               -> running
}

type link struct {
	mu     sync.Mutex
	cond   *sync.Cond
	result any
}

func makeLink() *link {
	l := &link{}
	l.cond = sync.NewCond(&l.mu)
	return l
}

type defaultSchedulerLinker struct {
	gid            int64
	s              Scheduler
	linkId2LinkMap map[string]*link
}

func makeDefaultSchedulerLinker(gid int64, s Scheduler) *defaultSchedulerLinker {
	sl := &defaultSchedulerLinker{}
	sl.gid = gid
	sl.linkId2LinkMap = make(map[string]*link)
	sl.s = s
	return sl
}

func (sl *defaultSchedulerLinker) GoroutineId() int64 {
	return sl.gid
}

func (sl *defaultSchedulerLinker) LinkAndWait(linkId string) any {
	if sl.gid == sl.s.GoroutineId() {
		if _, ok := sl.linkId2LinkMap[linkId]; ok {
			panic(fmt.Sprintf("create link error: linkId %s alreay exists", linkId))
		}
		sl.linkId2LinkMap[linkId] = makeLink()
	}

	link, ok := sl.linkId2LinkMap[linkId]
	if !ok {
		panic(fmt.Sprintf("reuse link error: linkId %s not found", linkId))
	}

	link.mu.Lock()
	defer link.mu.Unlock()

	if sl.gid == sl.s.GoroutineId() {
		nsl := &defaultSchedulerLinker{
			gid:            sl.gid + 1,
			s:              sl.s,
			linkId2LinkMap: sl.linkId2LinkMap,
		}
		sl.s.start(nsl)
	} else {
		link.cond.Signal()
	}
	link.cond.Wait()
	result := link.result
	return result
}

func (sl *defaultSchedulerLinker) SetResultAndWait(linkId string, result any) {
	if sl.gid != sl.s.GoroutineId() {
		panic("SetResultAndWait error: goroutine is not scheduling goroutine")
	}

	link, ok := sl.linkId2LinkMap[linkId]
	if !ok {
		panic(fmt.Sprintf("SetResultAndWait error: linkId %s not found", linkId))
	}

	link.mu.Lock()
	defer link.mu.Unlock()

	link.result = result
	link.cond.Signal()
	link.cond.Wait()
}

func (sl *defaultSchedulerLinker) Unlink(linkId string) {
	if sl.gid == sl.s.GoroutineId() {
		panic("Unlink error: goroutine is scheduling goroutine")
	}

	link, ok := sl.linkId2LinkMap[linkId]
	if !ok {
		panic(fmt.Sprintf("Unlink error: linkId %s not found", linkId))
	}

	link.mu.Lock()
	defer link.mu.Unlock()

	link.cond.Signal()
	delete(sl.linkId2LinkMap, linkId)
}

type Scheduler interface {
	// id of goroutine scheduling the stage.
	GoroutineId() int64
	// schedule the stages in an indefinite loop, can only be called once.
	Start(...SchedulerStage)
	// schedule stages.
	start(SchedulerLinker)
	// pause sheduling loop if it is running
	Pause()
	// resume scheduling loop if it is paused
	Resume()
	// stop scheduling loop
	Stop()
	// wait stop
	WaitStopped()
}

type schedulerState int64

const (
	New schedulerState = iota
	Running
	Paused
	Stopped
)

type defaultScheduler struct {
	mu         sync.Mutex
	cond       *sync.Cond
	gid        int64
	state      schedulerState
	stages     []SchedulerStage
	stageIndex int
}

func MakeDefaultScheduler() Scheduler {
	s := &defaultScheduler{}
	s.cond = sync.NewCond(&s.mu)
	return s
}

func (s *defaultScheduler) GoroutineId() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.gid
}

func (s *defaultScheduler) Start(stages ...SchedulerStage) {
	s.mu.Lock()
	if s.state != New {
		s.mu.Unlock()
		panic("Start error: Scheduler has started")
	}
	s.state = Running
	s.stages = stages
	s.mu.Unlock()

	s.start(makeDefaultSchedulerLinker(1, s))
}

func (s *defaultScheduler) start(sl SchedulerLinker) {
	s.mu.Lock()
	if s.gid+1 != sl.GoroutineId() {
		s.mu.Unlock()
		panic("start error: only scheduling goroutine can fork new goroutine")
	}
	s.gid = sl.GoroutineId()
	s.mu.Unlock()

	go s.schedule(sl)
}

func (s *defaultScheduler) schedule(sl SchedulerLinker) {
	for {
		s.mu.Lock()
		if s.gid != sl.GoroutineId() {
			s.mu.Unlock()
			break
		}
		for s.state == Paused {
			s.cond.Wait()
		}
		if s.state == Stopped {
			s.mu.Unlock()
			break
		}
		stage := s.stages[s.stageIndex]
		s.mu.Unlock()

		stage.Run(sl)
	}
}

func (s *defaultScheduler) Pause() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state == New {
		panic("Pause error: Scheduler has not started")
	}
	if s.state == Running {
		s.state = Paused
	}
}

func (s *defaultScheduler) Resume() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state == New {
		panic("Resume error: Scheduler has not started")
	}
	if s.state == Paused {
		s.state = Running
		s.cond.Broadcast()
	}
}

func (s *defaultScheduler) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.state == New {
		panic("Stop error: Scheduler has not started")
	}
	s.state = Stopped
	s.cond.Broadcast()
}

func (s *defaultScheduler) WaitStopped() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.state != Stopped {
		s.cond.Wait()
	}
}
