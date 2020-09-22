package migration

type Context struct {
	operations []DBOperation
}

func NewContext() *Context{
	return &Context{operations: make([]DBOperation, 0)}
}

func (ctx *Context) AddOperation(op DBOperation) {
	ctx.operations = append(ctx.operations, op)
}

func (ctx *Context) GetAllOperations() []DBOperation {
	return ctx.operations
}
