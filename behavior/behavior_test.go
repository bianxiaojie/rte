package behavior

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bianxiaojie/rte/utils/ref"
)

type testBehaviorObject struct {
}

func (o *testBehaviorObject) F(s string) string {
	return s
}

func (o *testBehaviorObject) F_1() {
}

func (o *testBehaviorObject) F_2(s string) string {
	return s
}

func (o *testBehaviorObject) F_3(s string, i int) string {
	return fmt.Sprintf("%s%d", s, i)
}

func (o *testBehaviorObject) F_1_1(s string) string {
	return s
}

func TestParseBehavior(t *testing.T) {
	ot := ref.ParseType[*testBehaviorObject]()
	f, _ := ot.MethodByName("F")
	f1, _ := ot.MethodByName("F_1")
	f2, _ := ot.MethodByName("F_2")
	f3, _ := ot.MethodByName("F_3")
	f11, _ := ot.MethodByName("F_1_1")

	parser := MakeDefaultBehaviorParser()

	_, ok := parser.ParseBehavior(f)
	if ok {
		t.Fatal("parse F: expected false, got true")
	}

	_, ok = parser.ParseBehavior(f1)
	if !ok {
		t.Fatal("parse F_1: expected true, got false")
	}

	_, ok = parser.ParseBehavior(f2)
	if !ok {
		t.Fatal("parse F_2: expected true, got false")
	}

	_, ok = parser.ParseBehavior(f3)
	if !ok {
		t.Fatal("parse F_3: expected true, got false")
	}

	_, ok = parser.ParseBehavior(f11)
	if ok {
		t.Fatal("parse F_1_1: expected false, got true")
	}
}

func TestCallBehavior(t *testing.T) {
	o := &testBehaviorObject{}
	ot := reflect.TypeOf(o)
	f1, _ := ot.MethodByName("F_1")
	f2, _ := ot.MethodByName("F_2")
	f3, _ := ot.MethodByName("F_3")

	parser := MakeDefaultBehaviorParser()

	b1, _ := parser.ParseBehavior(f1)
	result := b1.Call(o)
	if len(result) != 0 {
		t.Fatal("call F_1: expected none return value, got", result)
	}

	b2, _ := parser.ParseBehavior(f2)
	result = b2.Call(o, "s")
	if len(result) != 1 || result[0] != "s" {
		t.Fatal("call F_2: expected s, got", result)
	}

	b3, _ := parser.ParseBehavior(f3)
	result = b3.Call(o, "s", 1)
	if len(result) != 1 || result[0] != "s1" {
		t.Fatal("call F_3: expected s1, got", result)
	}
}
