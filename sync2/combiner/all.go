package combiner

var All = Descs{
	{"Mutex", false, func(bat Batcher, bound int) Combiner { return NewMutex(bat) }},
	{"Basic", false, func(bat Batcher, bound int) Combiner { return NewBasic(bat) }},
	{"BasicU", false, func(bat Batcher, bound int) Combiner { return NewBasicU(bat) }},
	{"Bounded", true, func(bat Batcher, bound int) Combiner { return NewBounded(bat, bound) }},
	{"BoundedU", true, func(bat Batcher, bound int) Combiner { return NewBoundedU(bat, bound) }},
	{"TBB", false, func(bat Batcher, bound int) Combiner { return NewTBB(bat) }},
}
