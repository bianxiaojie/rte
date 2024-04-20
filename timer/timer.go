package timer

import "time"

type Timer interface {
	// get current time
	GetTime() time.Duration
	// get time unit
	GetTimeunit() time.Duration
	// get terminal time
	GetTerminalTime() time.Duration
	// set terminal time
	SetTerminalTime(time.Duration)
}

type IncrementalTimer interface {
	Timer
	Increment()
}

type defaultIncrementalTimer struct {
	count        uint64
	unit         time.Duration
	terminalTime time.Duration
}

func MakeDefaultIncrementalTimer(unit time.Duration, terminalTime time.Duration) IncrementalTimer {
	it := &defaultIncrementalTimer{}
	it.unit = unit
	it.terminalTime = terminalTime
	return it
}

func (it *defaultIncrementalTimer) GetTime() time.Duration {
	return time.Duration(it.count) * it.unit
}

func (it *defaultIncrementalTimer) GetTimeunit() time.Duration {
	return it.unit
}

func (it *defaultIncrementalTimer) GetTerminalTime() time.Duration {
	return it.terminalTime
}

func (it *defaultIncrementalTimer) SetTerminalTime(terminalTime time.Duration) {
	it.terminalTime = terminalTime
}

func (it *defaultIncrementalTimer) Increment() {
	it.count++
}
