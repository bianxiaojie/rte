package entity

import (
	"testing"

	"github.com/bianxiaojie/rte/utils/ref"
	"github.com/bianxiaojie/rte/utils/sli"
)

type testContainerEntity struct {
	id string
}

func (e *testContainerEntity) Id() string {
	return e.id
}

func TestEntityContainer(t *testing.T) {
	e1 := &testContainerEntity{id: "e1"}
	e2 := &testContainerEntity{id: "e2"}
	e3 := &testContainerEntity{id: "e3"}

	ec := MakeDefaultEntityContainer()
	ec.AddEntity(e1)
	if _, ok := ec.GetEntityById("e1"); !ok {
		t.Fatal("container: fail to get e1")
	}

	ec.RemoveEntityById("e1")
	if _, ok := ec.GetEntityById("e1"); ok {
		t.Fatal("container: fail to remove e1")
	}

	ec.AddEntity(e1)
	ec.AddEntity(e2)
	ec.AddEntity(e3)
	entities := ec.GetEntitiesByType(ref.ParseType[*testContainerEntity]())
	if !sli.EqualIgnoreOrder([]Entity{e1, e2, e3}, entities) {
		t.Fatalf("container: expected %v, got %v\n", []Entity{e1, e2, e3}, entities)
	}

	entities = ec.GetEntitiesByIdPattern(`e\d+`)
	if !sli.EqualIgnoreOrder([]Entity{e1, e2, e3}, entities) {
		t.Fatalf("container: expected %v, got %v\n", []Entity{e1, e2, e3}, entities)
	}

	entities = ec.GetEntities()
	if !sli.EqualIgnoreOrder([]Entity{e1, e2, e3}, entities) {
		t.Fatalf("container: expected %v, got %v\n", []Entity{e1, e2, e3}, entities)
	}

	ec.RemoveEntitiesByIdPattern(`e1`)
	entities = ec.GetEntities()
	if !sli.EqualIgnoreOrder([]Entity{e2, e3}, entities) {
		t.Fatalf("container: expected %v, got %v\n", []Entity{e2, e3}, entities)
	}

	ec.RemoveEntitiesByType(ref.ParseType[testContainerEntity]())
	entities = ec.GetEntities()
	if !sli.EqualIgnoreOrder([]Entity{e2, e3}, entities) {
		t.Fatalf("container: expected %v, got %v\n", []Entity{e2, e3}, entities)
	}

	ec.RemoveEntitiesByType(ref.ParseType[*testContainerEntity]())
	entities = ec.GetEntities()
	if len(entities) != 0 {
		t.Fatalf("container: expected entities empty, got %v\n", entities)
	}
}
