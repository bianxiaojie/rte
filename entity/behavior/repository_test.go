package behavior

import (
	"testing"

	"github.com/bianxiaojie/rte/utils/ref"
	"github.com/bianxiaojie/rte/utils/sli"
)

type testRepositoryObject1 struct {
}

func (o *testRepositoryObject1) F1_1() {

}

func (o *testRepositoryObject1) F2_3() {

}

type testRepositoryObject2 struct {
}

func (o *testRepositoryObject2) G1_1() {

}

func (o *testRepositoryObject2) G2_3() {

}

func TestBehaviorRepository(t *testing.T) {
	br := MakeDefaultBehaviorRepository()
	parser := MakeDefaultBehaviorParser()

	sortedBehaviors := br.GetSortedBehaviors()
	if len(sortedBehaviors) != 0 {
		t.Fatal("repository: expected len 0, got", len(sortedBehaviors))
	}

	o1t := ref.ParseType[*testRepositoryObject1]()
	f1, _ := o1t.MethodByName("F1_1")
	bf1, _ := parser.ParseBehavior(f1)
	f2, _ := o1t.MethodByName("F2_3")
	bf2, _ := parser.ParseBehavior(f2)
	o2t := ref.ParseType[*testRepositoryObject2]()
	g1, _ := o2t.MethodByName("G1_1")
	bg1, _ := parser.ParseBehavior(g1)
	g2, _ := o2t.MethodByName("G2_3")
	bg2, _ := parser.ParseBehavior(g2)

	br.AddBehavior(bf1)
	sortedBehaviors = br.GetSortedBehaviors()
	if !sortedBehaviorsEqual([][]Behavior{{bf1}}, sortedBehaviors) {
		t.Fatalf("repository: expected %v, got %v\n", [][]Behavior{{bf1}}, sortedBehaviors)
	}

	br.AddBehaviorByType(o1t)
	sortedBehaviors = br.GetSortedBehaviors()
	if !sortedBehaviorsEqual([][]Behavior{{bf1}, {bf2}}, sortedBehaviors) {
		t.Fatalf("repository: expected %v, got %v\n", [][]Behavior{{bf1}, {bf2}}, sortedBehaviors)
	}

	br.RemoveBehavior(bf1)
	sortedBehaviors = br.GetSortedBehaviors()
	if !sortedBehaviorsEqual([][]Behavior{{bf2}}, sortedBehaviors) {
		t.Fatalf("repository: expected %v, got %v\n", [][]Behavior{{bf2}}, sortedBehaviors)
	}

	br.RemoveBehaviorByType(o1t)
	sortedBehaviors = br.GetSortedBehaviors()
	if len(sortedBehaviors) != 0 {
		t.Fatal("repository: expected len 0, got", len(sortedBehaviors))
	}

	br.AddBehaviorByType(o1t)
	br.AddBehaviorByType(o2t)
	sortedBehaviors = br.GetSortedBehaviors()
	if !sortedBehaviorsEqual([][]Behavior{{bf1, bg1}, {bf2, bg2}}, sortedBehaviors) {
		t.Fatalf("repository: expected %v, got %v\n", [][]Behavior{{bf1, bg1}, {bf2, bg2}}, sortedBehaviors)
	}
}

func sortedBehaviorsEqual(expected [][]Behavior, actual [][]Behavior) bool {
	if len(expected) != len(actual) {
		return false
	}

	for i := 0; i < len(expected); i++ {
		if !sli.EqualIgnoreOrderFunc(expected[i], actual[i], Behavior.Equal) {
			return false
		}
	}

	return true
}
