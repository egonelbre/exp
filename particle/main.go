package main

import (
	"fmt"
	"runtime"
	"unsafe"
)

func main() {
	particles := Particles{}
	defer runtime.KeepAlive(particles)

	for i := 0; i < ParticleCount; i++ {
		p := particles.Ref(i)
		p.X = float32(i)
		p.Y = float32(-i)
		p.Z = float32(i)
	}

	fmt.Println(particles.X)
	fmt.Println(particles.Y)
	fmt.Println(particles.Z)
}

const ParticleCount = 256

type Particles struct {
	X [ParticleCount]float32
	Y [ParticleCount]float32
	Z [ParticleCount]float32

	VX [ParticleCount]float32
	VY [ParticleCount]float32
	VZ [ParticleCount]float32

	CA [ParticleCount]float32
	CR [ParticleCount]float32
	CG [ParticleCount]float32
	CB [ParticleCount]float32
}

func (p *Particles) Ref(i int) *Particle {
	return (*Particle)(unsafe.Pointer(&p.X[i]))
}

type Particle struct {
	X float32
	_ [ParticleCount - 1]float32
	Y float32
	_ [ParticleCount - 1]float32
	Z float32
	_ [ParticleCount - 1]float32

	VX float32
	_  [ParticleCount - 1]float32
	VY float32
	_  [ParticleCount - 1]float32
	VZ float32

	CA float32
	_  [ParticleCount - 1]float32
	CR float32
	_  [ParticleCount - 1]float32
	CG float32
	_  [ParticleCount - 1]float32
	CB float32
}
