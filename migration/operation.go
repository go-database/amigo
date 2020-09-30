package migration

type (
	Operation interface {
	}
	OperationList []Operation
)

func (l *OperationList) Add(op Operation){
	*l = append(*l, op)
}

func (l *OperationList) Count() int{
	return len(*l)
}
