package migration

import "testing"

func TestOperationList_Add(t *testing.T) {
	l := OperationList{}
	l.Add(&EmptyOperation{})
	if len(l) == 0{
		t.Error()
	}
}

func TestOperationList_Count(t *testing.T) {
	l := OperationList{}
	l.Add(&EmptyOperation{})
	if l.Count() != 1{
		t.Error()
	}
}
