package combiner

type Lock struct {
	exe Executor
}

func NewLock(exe Executor) *Lock { return &Lock{exe} }

func (lock *Lock) Do(v int64) {
	lock.exe.Start()
	lock.exe.Include(v)
	lock.exe.Finish()
}
