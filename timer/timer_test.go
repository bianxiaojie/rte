package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	it := MakeDefaultIncrementalTimer(time.Second, time.Second)

	it.Increment()
	if time.Second != it.GetTime() {
		t.Fatalf("timer increment error: expected %d, got %d\n", time.Second, it.GetTime())
	}
}
