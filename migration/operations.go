package migration

type (
	EmptyDBOperation struct {
	}
	TableOperation struct {
		TableName string
		Schema string
	}
	DropTableOperation struct {
		TableOperation
	}
)

var (
	_ DBOperation = &EmptyDBOperation{}
	_ DBOperation = &TableOperation{}
)
