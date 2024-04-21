package engine

import (
	"fmt"
	"slices"
	"testing"
	"time"

	"github.com/bianxiaojie/rte/ctx"
	"github.com/bianxiaojie/rte/entity"
	"github.com/bianxiaojie/rte/entity/action"
	"github.com/bianxiaojie/rte/utils/ref"
)

type testEngineEntity struct {
	result *[]string
	id     string
}

func (e *testEngineEntity) Id() string {
	return e.id
}

func (e *testEngineEntity) AppendResult(result string) {
	*e.result = append(*e.result, result)
}

func (e *testEngineEntity) Count_1(context ctx.Context) {
	action.HandleNoneTargetAction[*countAction, countEntity](context.ActionHandler(), e, nil)
}

func (e *testEngineEntity) Heartbeat_2(context ctx.Context) {
	r := action.HandleNoneTargetStagedAction[*heartbeatAction, *heartbeatStage, heartbeatEntity](context.ActionHandler(), e, 3)
	*e.result = append(*e.result, r)
}

// definition for count action
type countEntity interface {
	entity.Entity
	AppendResult(string)
}

type countAction struct {
}

func (ca *countAction) Action(ce countEntity, param any) any {
	ce.AppendResult(fmt.Sprintf("%s Count", ce.Id()))
	return nil
}

// definition for heartbeat action
type heartbeatStage struct {
	count int
	value string
}

func (hbs *heartbeatStage) IsLastStage() bool {
	return hbs.count <= 0
}

func (hbs *heartbeatStage) GetReturnedValue() string {
	return hbs.value
}

type heartbeatEntity interface {
	entity.Entity
	AppendResult(string)
}

type heartbeatAction struct {
}

func (hba *heartbeatAction) MakeStage(hbe heartbeatEntity, count int) *heartbeatStage {
	hbs := &heartbeatStage{}
	hbs.count = count
	return hbs
}

func (hba *heartbeatAction) ActionStage(hbe heartbeatEntity, count int, hbs *heartbeatStage) *heartbeatStage {
	hbe.AppendResult(fmt.Sprintf("%s Heartbeat", hbe.Id()))
	hbs.value = fmt.Sprintf("%s%s", hbs.value, hbe.Id())
	hbs.count--
	return hbs
}

func TestEngine(t *testing.T) {
	expected := make([]string, 0)
	for i := 1; i <= 10; i++ {
		expected = append(expected, "e1 Count")
		expected = append(expected, "e2 Count")
		expected = append(expected, "e1 Heartbeat")
		if i%3 == 0 {
			expected = append(expected, "e1e1e1")
		}
		expected = append(expected, "e2 Heartbeat")
		if i%3 == 0 {
			expected = append(expected, "e2e2e2")
		}
	}
	result := make([]string, 0)

	e := MakeDefaultEngine(time.Second, 10*time.Second)
	em := e.EntityManager()
	em.AddBehaviorByType(ref.ParseType[*testEngineEntity]())
	em.AddEntity(&testEngineEntity{id: "e1", result: &result})
	em.AddEntity(&testEngineEntity{id: "e2", result: &result})
	e.Start()
	e.WaitStopped()
	if !slices.Equal(expected, result) {
		t.Fatalf("test engine error: expected: %v, got %v\n", expected, result)
	}
}
