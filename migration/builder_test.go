package migration

import (
	"testing"
)

func TestBuilder_BuildDropTable(t *testing.T) {
	const name = "test"
	c := NewContext()
	b := Builder{ctx: c}
	b.DropTable(name).InSchema(name)
	op, ok := c.operations[0].(*DropTableOperation)
	if !ok {
		t.Error("operation not found")
	}
	if op.TableName != name || op.Schema != name {
		t.Error()
	}
}