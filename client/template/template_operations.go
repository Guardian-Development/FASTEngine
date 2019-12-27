package template

const constantOperation = "constant"

type Operation interface {
}

type OperationNone struct {
}

type OperationConstant struct {
	constantValue string
}
