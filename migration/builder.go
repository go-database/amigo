package migration

type Builder struct {
	ctx *Context
}

func (b *Builder) CreateTable(tableName string) *CreateTableOperationBuilder{
	op := &CreateTableOperation{}
	op.TableName = tableName
	b.ctx.AddOperation(op)
	return &CreateTableOperationBuilder{op}
}

func (b *Builder) DropTable(tableName string) *DropTableOperationBuilder {
	op := &DropTableOperation{}
	op.TableName = tableName
	b.ctx.AddOperation(op)
	return &DropTableOperationBuilder{op: op}
}