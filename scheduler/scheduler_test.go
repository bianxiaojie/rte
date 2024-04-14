package scheduler

import (
	"slices"
	"testing"
)

type testSchedulerStage struct {
	counter      uint
	maxCount     uint
	stepCounter  uint
	maxStepCount uint
	result       []uint
	linkId       string
	s            Scheduler
}

func makeTestSchedulerStage(maxCount uint, maxStepCount uint, s Scheduler) *testSchedulerStage {
	ss := &testSchedulerStage{}
	ss.maxCount = maxCount
	ss.maxStepCount = maxStepCount
	ss.result = make([]uint, 0)
	ss.s = s
	return ss
}

func (ss *testSchedulerStage) Run(sl SchedulerLinker) {
	if ss.counter >= ss.maxCount {
		ss.s.Stop()
		return
	}

	linkId := "linkId"
	if len(ss.linkId) > 0 {
		sl.SetResultAndWait(linkId, nil)
	}

	if ss.stepCounter >= ss.maxStepCount {
		ss.stepCounter = 0
		ss.linkId = linkId
		sl.LinkAndWait(linkId)
		ss.linkId = ""
		sl.Unlink(linkId)
	} else {
		ss.result = append(ss.result, ss.counter)
		ss.counter++
		ss.stepCounter++
	}
}

func TestScheduler(t *testing.T) {
	s := MakeDefaultScheduler()
	ss := makeTestSchedulerStage(100, 10, s)
	s.Start(ss)
	s.WaitStopped()

	result := make([]uint, 100)
	for i := 0; i < 100; i++ {
		result[i] = uint(i)
	}
	if !slices.Equal(result, ss.result) {
		t.Fatalf("expected %v, got %v\n", result, ss.result)
	}
}
