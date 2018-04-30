package queue

import "github.com/tidwall/fastlane"

// MPMC wrapper

type MPMCnw_fl struct {
	ch fastlane.ChanUint64
}

func NewMPMCnw_fl() *MPMCnw_fl {
	return &MPMCnw_fl{}
}

func (q *MPMCnw_fl) MultipleConsumers() {}
func (q *MPMCnw_fl) MultipleProducers() {}

func (q *MPMCnw_fl) Send(v Value) bool {
	q.ch.Send(uint64(v))
	return true
}

func (q *MPMCnw_fl) Recv(v *Value) bool {
	*v = Value(q.ch.Recv())
	return true
}
