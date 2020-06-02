package goroutine_pool

type Task struct {
	f func() error // 一个Task里面应该有一个具体的业务，业务的名称为f
}

// 创建一个任务
func NewTask(argF func() error) *Task {
	return &Task{f: argF}
}

// Execute 执行任务
func (t *Task) Execute() {
	_ = t.f()
}