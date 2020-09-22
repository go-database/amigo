package migration

type (
	EmptyDBOperation struct {
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
	_ DBOperation = &EmptyDBOperation{}
	_ DBOperation = &TableOperation{}
	_ DBOperation = &CreateTableOperation{}
	_ DBOperation = &DropTableOperation{}
)
