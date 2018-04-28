package queue

type Ctor func(size int) Queue

func (ctor Ctor) BlockingSPSC() func(size int) BlockingSPSC {
	if _, ok := ctor(1).(BlockingSPSC); ok {
		return func(size int) BlockingSPSC { return ctor(size).(BlockingSPSC) }
	}
	return nil
}

func (ctor Ctor) BlockingSPMC() func(size int) BlockingSPMC {
	if _, ok := ctor(1).(BlockingSPMC); ok {
		return func(size int) BlockingSPMC {
			return ctor(size).(BlockingSPMC)
		}
	}
	return nil
}

func (ctor Ctor) BlockingMPSC() func(size int) BlockingMPSC {
	if _, ok := ctor(1).(BlockingMPSC); ok {
		return func(size int) BlockingMPSC {
			return ctor(size).(BlockingMPSC)
		}
	}
	return nil
}

func (ctor Ctor) BlockingMPMC() func(size int) BlockingMPMC {
	if _, ok := ctor(1).(BlockingMPMC); ok {
		return func(size int) BlockingMPMC {
			return ctor(size).(BlockingMPMC)
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
