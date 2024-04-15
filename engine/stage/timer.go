package stage

import (
	"github.com/bianxiaojie/rte/scheduler"
	"github.com/bianxiaojie/rte/timer"
)

type StageTimer struct {
	s scheduler.Scheduler
	t timer.IncrementalTimer
}

func MakeStageTimer(s scheduler.Scheduler, t timer.IncrementalTimer) *StageTimer {
	st := &StageTimer{}
	st.s = s
	st.t = t
	return st
}

func (st *StageTimer) Run(sl scheduler.SchedulerLinker) {
	if st.t.GetTime() >= st.t.GetTerminalTime() {
		st.s.Stop()
		return
	}

	st.t.Increment()
}
