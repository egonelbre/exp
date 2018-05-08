package combiner

var All = Descs{
	{"Mutex", false, func(bat Batcher, bound int) Combiner { return NewMutex(bat) }},
	{"SpinMutex", false, func(bat Batcher, bound int) Combiner { return NewSpinMutex(bat) }},
	{"Basic", false, func(bat Batcher, bound int) Combiner { return NewBasic(bat) }},
	{"BasicS", false, func(bat Batcher, bound int) Combiner { return NewBasicSleepy(bat) }},
	{"BasicU", false, func(bat Batcher, bound int) Combiner { return NewBasicUintptr(bat) }},
	{"BasicSU", false, func(bat Batcher, bound int) Combiner { return NewBasicSleepyUintptr(bat) }},
	{"Bounded", true, func(bat Batcher, bound int) Combiner { return NewBounded(bat, bound) }},
	{"BoundedS", true, func(bat Batcher, bound int) Combiner { return NewBoundedSleepy(bat, bound) }},
	{"BoundedU", true, func(bat Batcher, bound int) Combiner { return NewBoundedUintptr(bat, bound) }},
	{"BoundedSU", true, func(bat Batcher, bound int) Combiner { return NewBoundedSleepyUintptr(bat, bound) }},
	{"TBB", false, func(bat Batcher, bound int) Combiner { return NewTBB(bat) }},
	{"TBBS", false, func(bat Batcher, bound int) Combiner { return NewTBBSleepy(bat) }},
	{"TBBU", false, func(bat Batcher, bound int) Combiner { return NewTBBUintptr(bat) }},
	{"TBBSU", false, func(bat Batcher, bound int) Combiner { return NewTBBSleepyUintptr(bat) }},
	{"Remote", false, func(bat Batcher, bound int) Combiner { return NewRemote(bat) }},
}
