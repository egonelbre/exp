package queue

import "github.com/tidwall/fastlane"

// MPSC wrapper

type MPSCnw_fl struct {
	ch fastlane.ChanUint64
}

func NewMPSCnw_fl() *MPSCnw_fl {
	return &MPSCnw_fl{}
}

func (q *MPSCnw_fl) MultipleProducers() {}

func (q *MPSCnw_fl) Send(v Value) bool {
	q.ch.Send(uint64(v))
	return true
}

func (q *MPSCnw_fl) Recv(v *Value) bool {
	*v = Value(q.ch.Recv())
	return true
}
