package combiner

var All = Descs{
	{"Nop", false, func(bat Batcher, bound int) Combiner { return NewNop(bat) }},
	{"Basic", false, func(bat Batcher, bound int) Combiner { return NewBasic(bat) }},
	{"Bounded", true, func(bat Batcher, bound int) Combiner { return NewBounded(bat, bound) }},
	{"BoundedU", true, func(bat Batcher, bound int) Combiner { return NewBoundedU(bat, bound) }},
	{"TBB", false, func(bat Batcher, bound int) Combiner { return NewTBB(bat) }},
}
