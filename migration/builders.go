package migration

type (
	DropTableOperationBuilder struct {
		op *DropTableOperation
	}
)

func (b *DropTableOperationBuilder) InSchema(schema string) {
	b.op.Schema = schema
}