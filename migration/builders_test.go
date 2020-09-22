package migration

import "testing"

func TestDropTableOperationBuilder_InSchema(t *testing.T) {
	const schemaName = "test"
	op := &DropTableOperation{}
	b := &DropTableOperationBuilder{
		op: op,
	}
	b.InSchema(schemaName)
	if op.Schema != schemaName {
		t.Error()
	}
}