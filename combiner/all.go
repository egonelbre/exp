package combiner

import "github.com/loov/combiner/testsuite"

var All = testsuite.Descs{
	{"TBB", false, func(bat testsuite.Batcher, bound int) testsuite.Combiner { return NewTBB(bat) }},
	{"TBBS", false, func(bat testsuite.Batcher, bound int) testsuite.Combiner { return NewTBBSleepy(bat) }},
	{"TBBU", false, func(bat testsuite.Batcher, bound int) testsuite.Combiner { return NewTBBUintptr(bat) }},
	{"TBBSU", false, func(bat testsuite.Batcher, bound int) testsuite.Combiner { return NewTBBSleepyUintptr(bat) }},
	{"Remote", false, func(bat testsuite.Batcher, bound int) testsuite.Combiner { return NewRemote(bat) }},
}
