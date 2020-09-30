package migration

type (
	EmptyOperation struct {
	}
	TableOperation struct {
		TableName string
		Schema    string
	}
	CreateTableOperation struct {
		TableOperation
	}
	DropTableOperation struct {
		TableOperation
	}
)

var (
	_ Operation = &EmptyOperation{}
	_ Operation = &TableOperation{}
	_ Operation = &CreateTableOperation{}
	_ Operation = &DropTableOperation{}
)
