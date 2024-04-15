package action

import "testing"

type testActionStage struct {
}

func (as *testActionStage) IsLastStage() bool {
	return true
}

func (as *testActionStage) GetReturnedValue() any {
	return nil
}

type testActionObject struct {
	s string
}

func (ao *testActionObject) Action(s string, ignore any) any {
	return nil
}

func (ao *testActionObject) MakeStage(p any) *testActionStage {
	return &testActionStage{}
}

func (ao *testActionObject) ActionStage(s string, as *testActionStage) *testActionStage {
	return as
}

func TestAction(t *testing.T) {
	af := MakeDefaultActionFactory()
	nta := GetNoneTargetAction[*testActionObject](af)
	ntsa := GetNoneTargetStagedAction[*testActionObject](af)
	if nta != ntsa {
		t.Fatalf("get action: expected %p == %p, got false", nta, ntsa)
	}

	af = MakeDefaultActionFactoryWithParams(false)
	nta = GetNoneTargetAction[*testActionObject](af)
	ntsa = GetNoneTargetStagedAction[*testActionObject](af)
	if nta == ntsa {
		t.Fatalf("get action: expected %p != %p, got true", nta, ntsa)
	}
}
