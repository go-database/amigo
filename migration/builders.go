package migration

type (
	CreateTableOperationBuilder struct {
		op *CreateTableOperation
	}
	DropTableOperationBuilder struct {
		op *DropTableOperation
	}
)

func (b *CreateTableOperationBuilder) InSchema(schema string) {
	b.op.Schema = schema
}

func (b *DropTableOperationBuilder) InSchema(schema string) {
	b.op.Schema = schema
}