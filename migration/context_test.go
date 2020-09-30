package migration

import "testing"

func Test_NewContext(t *testing.T) {
	ctx := NewContext()
	if ctx.operations == nil {
		t.Error("")
	}
}

func TestContext_GetAllOperations(t *testing.T) {
	ctx := NewContext()
	ctx.operations.Add(&EmptyOperation{})
	if ctx.GetAllOperations() == nil {
		t.Error("operations is nil")
	}
	if len(ctx.GetAllOperations()) == 0 {
		t.Error()
	}
}

func TestContext_SetDefaultSchema(t *testing.T) {
	ctx := NewContext()
	ctx.SetDefaultSchema("test")
	if ctx.defaultSchema != "test" {
		t.Error()
	}
}

func TestContext_GetDefaultSchema(t *testing.T) {
	ctx := &Context{
		defaultSchema: "test",
	}
	if ctx.GetDefaultSchema() != "test" {
		t.Error()
	}
}