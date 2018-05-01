package queue

import (
	"github.com/loov/queue/extqueue"
	"github.com/loov/queue/testsuite"
)

var Descs = []extqueue.Desc{
	{"SPSCrMC", extqueue.Batched, func(bs, s int) testsuite.Queue { return NewSPSCrMC(bs, s) }},
	// {"SPSCrsMC", extqueue.Batched, func(bs, s int) testsuite.Queue { return NewSPSCrsMC(bs, s) }},
	// {"MPSCrMC", extqueue.Batched, func(bs, s int) testsuite.Queue { return NewMPSCrMC(bs, s) }},
	// {"MPSCrsMC", extqueue.Batched, func(bs, s int) testsuite.Queue { return NewMPSCrsMC(bs, s) }},
}
