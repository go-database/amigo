package migration

type Context struct {
	defaultSchema string
	operations []Operation
}

func NewContext() *Context{
	return &Context{operations: make([]Operation, 0)}
}

func (ctx *Context) AddOperation(op Operation) {
	ctx.operations = append(ctx.operations, op)
}

func (ctx *Context) GetAllOperations() []Operation {
	return ctx.operations
}

func (ctx *Context) SetDefaultSchema(name string){
	ctx.defaultSchema = name
}

func (ctx *Context) GetDefaultSchema() string{
	return ctx.defaultSchema
}
