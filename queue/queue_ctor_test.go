package queue

type Ctor func(size int) Queue

func (ctor Ctor) SPSC() func(size int) SPSC {
	if _, ok := ctor(1).(SPSC); ok {
		return func(size int) SPSC { return ctor(size).(SPSC) }
	}
	return nil
}

func (ctor Ctor) SPMC() func(size int) SPMC {
	if _, ok := ctor(1).(SPMC); ok {
		return func(size int) SPMC {
			return ctor(size).(SPMC)
		}
	}
	return nil
}

func (ctor Ctor) MPSC() func(size int) MPSC {
	if _, ok := ctor(1).(MPSC); ok {
		return func(size int) MPSC {
			return ctor(size).(MPSC)
		}
	}
	return nil
}

func (ctor Ctor) MPMC() func(size int) MPMC {
	if _, ok := ctor(1).(MPMC); ok {
		return func(size int) MPMC {
			return ctor(size).(MPMC)
		}
	}
	return nil
}

func (ctor Ctor) NonblockingSPSC() func(size int) NonblockingSPSC {
	if _, ok := ctor(1).(NonblockingSPSC); ok {
		return func(size int) NonblockingSPSC {
			return ctor(size).(NonblockingSPSC)
		}
	}
	return nil
}

func (ctor Ctor) NonblockingSPMC() func(size int) NonblockingSPMC {
	if _, ok := ctor(1).(NonblockingSPMC); ok {
		return func(size int) NonblockingSPMC {
			return ctor(size).(NonblockingSPMC)
		}
	}
	return nil
}

func (ctor Ctor) NonblockingMPSC() func(size int) NonblockingMPSC {
	if _, ok := ctor(1).(NonblockingMPSC); ok {
		return func(size int) NonblockingMPSC {
			return ctor(size).(NonblockingMPSC)
		}
	}
	return nil
}

func (ctor Ctor) NonblockingMPMC() func(size int) NonblockingMPMC {
	if _, ok := ctor(1).(NonblockingMPMC); ok {
		return func(size int) NonblockingMPMC {
			return ctor(size).(NonblockingMPMC)
		}
	}
	return nil
}
