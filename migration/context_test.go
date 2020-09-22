package migration

import "testing"

func Test_NewContext(t *testing.T) {
	ctx := NewContext()
	if ctx.operations == nil {
		t.Error("")
	}
}

func TestContext_AddOperation(t *testing.T) {
	ctx := NewContext()
	ctx.AddOperation(&EmptyDBOperation{})
	if len(ctx.operations) == 0 {
		t.Error()
	}
}

func TestContext_GetAllOperations(t *testing.T) {
	ctx := NewContext()
	ctx.AddOperation(&EmptyDBOperation{})
	if len(ctx.GetAllOperations()) == 0 {
		t.Error()
	}
}
