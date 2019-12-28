package operation

type Operation interface {
}

type None struct {
}

type Constant struct {
	ConstantValue string
}
