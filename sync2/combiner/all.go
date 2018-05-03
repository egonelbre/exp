package combiner

var All = Descs{
	{"Mutex", false, func(bat Batcher, bound int) Combiner { return NewMutex(bat) }},
	{"SpinMutex", false, func(bat Batcher, bound int) Combiner { return NewSpinMutex(bat) }},
	{"Basic", false, func(bat Batcher, bound int) Combiner { return NewBasic(bat) }},
	{"BasicU", false, func(bat Batcher, bound int) Combiner { return NewBasicU(bat) }},
	{"Bounded", true, func(bat Batcher, bound int) Combiner { return NewBounded(bat, bound) }},
	{"BoundedU", true, func(bat Batcher, bound int) Combiner { return NewBoundedU(bat, bound) }},
	{"TBB", false, func(bat Batcher, bound int) Combiner { return NewTBB(bat) }},
	{"TBBU", false, func(bat Batcher, bound int) Combiner { return NewTBBU(bat) }},
	{"Remote", false, func(bat Batcher, bound int) Combiner { return NewRemote(bat) }},
}
