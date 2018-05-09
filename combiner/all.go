package combiner

import "github.com/loov/combiner/testsuite"

var All = testsuite.Descs{
	{"TBB", false, func(bat Batcher, bound int) Combiner { return NewTBB(bat) }},
	{"TBBS", false, func(bat Batcher, bound int) Combiner { return NewTBBSleepy(bat) }},
	{"TBBU", false, func(bat Batcher, bound int) Combiner { return NewTBBUintptr(bat) }},
	{"TBBSU", false, func(bat Batcher, bound int) Combiner { return NewTBBSleepyUintptr(bat) }},
	{"Remote", false, func(bat Batcher, bound int) Combiner { return NewRemote(bat) }},
}

type Batcher = testsuite.Batcher
type Combiner = testsuite.Combiner
type Argument = testsuite.Argument
