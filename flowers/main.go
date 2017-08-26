package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:     "Flowers!",
		Bounds:    pixel.R(0, 0, 480, 320),
		Resizable: true,
		VSync:     true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	camera := pixel.V(0, 0)
	branch := NewRoot()

	canvas := imdraw.New(nil)
	background := 0.0
	now := time.Now()
	for !win.Closed() {
		win.SetClosed(win.JustPressed(pixelgl.KeyEscape))

		next := time.Now()
		dt := next.Sub(now).Seconds()
		now = next
		if dt > 1.0/15.0 {
			dt = 1.0 / 15.0
		}

		camera = pixel.Lerp(camera, branch.Head(), 0.7*dt)
		background += dt

		branch.Update(dt)

		win.Clear(HSL{background, 0.6, 0.99})
		//win.Clear(HSL{TAU / 3, 0.66, 0.99})

		canvas.Clear()
		canvas.Reset()
		canvas.SetMatrix(pixel.IM.Moved(win.Bounds().Center()).Moved(camera.Scaled(-1)))
		{
			branch.Draw(canvas)
		}
		canvas.Draw(win)

		win.Update()
	}
}

const ANGLESNAP = TAU / 8

type Branch struct {
	Time float64

	PathLimit int
	Path      []pixel.Vec

	Thickness  float64
	Lightness  float64
	Accelerate float64
	Speed      float64
	Direction  float64
	Turn       float64
	Length     float64
	Travel     float64

	IsRoot         bool
	Life           int
	SpawnCountdown float64
	SpawnInterval  float64
	Branches       []*Branch
}

func randomsnap(min, max float64, snap float64) float64 {
	return min + math.Floor(rand.Float64()*(max-min)/snap)*snap
}

func random(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func NewRoot() *Branch {
	branch := &Branch{}
	branch.PathLimit = 300
	branch.Speed = 100.0
	branch.Direction = random(0, TAU)
	branch.IsRoot = true
	branch.Thickness = 8.0
	branch.Lightness = 0.7

	branch.SpawnInterval = 0.3
	branch.SpawnCountdown = branch.SpawnInterval

	branch.NextSegment()

	return branch
}

func NewBranch(root *Branch) *Branch {
	child := NewRoot()
	child.Thickness = root.Thickness * 0.3
	child.IsRoot = false
	child.Lightness = root.Lightness * 0.5
	child.Life = len(root.Path)
	child.Path = []pixel.Vec{root.Head()}
	child.Direction = root.Direction
	child.Turn = randomsnap(-TAU*3/4, TAU*3/4, ANGLESNAP)
	return child
}

func (branch *Branch) Head() pixel.Vec {
	if len(branch.Path) == 0 {
		return pixel.V(0, 0)
	}
	return branch.Path[len(branch.Path)-1]
}

func (branch *Branch) NextSegment() {
	branch.Accelerate += 1.0
	branch.Turn = randomsnap(-TAU/2, TAU/2, ANGLESNAP)
	branch.Length = randomsnap(100, 200, 25)
	branch.Travel = branch.Length
}

func (branch *Branch) IsDying() bool {
	return !branch.IsRoot && (branch.Life < 0)
}
func (branch *Branch) IsDead() bool {
	return !branch.IsRoot && (branch.Life < 0) && (len(branch.Path) == 0)
}

func (branch *Branch) Update(dt float64) {
	branch.Life--

	alive := branch.Branches[:0]
	for _, child := range branch.Branches {
		child.Update(dt)
		if !child.IsDead() {
			alive = append(alive, child)
		}
	}
	branch.Branches = alive

	if branch.Accelerate > 0 {
		branch.Accelerate -= dt * 5
		dt += Bounce(branch.Accelerate, 0, 1, 1) * dt
	}
	branch.Time += dt

	if branch.IsRoot {
		branch.SpawnCountdown -= dt
		if branch.SpawnCountdown < 0 {
			branch.SpawnCountdown = branch.SpawnInterval
			child := NewBranch(branch)
			branch.Branches = append(branch.Branches, child)
		}
	}

	if !branch.IsDying() {
		distance := branch.Speed * dt

		if branch.IsRoot || branch.Travel > 0 {
			sn, cs := math.Sincos(branch.Direction)
			branch.Direction += branch.Turn * distance / branch.Length
			dir := pixel.V(cs, sn).Scaled(distance)
			branch.Travel -= distance

			branch.Path = append(branch.Path, branch.Head().Add(dir))
			if len(branch.Path) > branch.PathLimit {
				copy(branch.Path, branch.Path[1:])
				branch.Path = branch.Path[:len(branch.Path)-1]
			}
		}

		if branch.IsRoot && branch.Travel <= 0 {
			branch.NextSegment()
		}
	} else {
		branch.Path = branch.Path[1:]
	}
}

func (branch *Branch) Draw(m *imdraw.IMDraw) {
	for _, child := range branch.Branches {
		child.Draw(m)
	}

	Line(m, branch.Path, func(i int, at pixel.Vec) (float64, color.Color) {
		p := float64(i) / float64(branch.PathLimit)
		radius := branch.Thickness * (math.Sin(p*4*TAU+branch.Time*6) + 5) / 5

		pp := float64(i) / float64(len(branch.Path))
		if pp > 0.85 {
			radius *= 1 - (pp-0.85)/0.3
		}
		if !branch.IsRoot {
			if pp < 0.05 {
				radius *= pp / 0.05
			}
		}
		color := HSL{branch.Time + p*TAU/8, 0.3, 1 - branch.Lightness}
		return radius, color
	})
}

func Line(m *imdraw.IMDraw, path []pixel.Vec, attrfn func(i int, at pixel.Vec) (float64, color.Color)) {
	if len(path) < 2 {
		return
	}

	a := path[0]
	var x1, x2, xn pixel.Vec

	for i, b := range path {
		if i > 0 && a.Sub(b).Len() < 1 {
			continue
		}
		radius, color := attrfn(i, b)
		m.Color = color

		abn := ScaleTo(SegmentNormal(a, b), radius)
		// segment-corners
		a1, a2 := a.Add(abn), a.Sub(abn)
		b1, b2 := b.Add(abn), b.Sub(abn)

		if i > 0 && radius > 0.5 {
			d := Rotate(xn).Dot(abn)
			if d < 0 {
				m.Push(x1, a1, a)
			} else {
				m.Push(x2, a2, a)
			}
			m.Polygon(0)
		}

		m.Push(a1, b1, b2, a2)
		m.Polygon(0)

		a = b
		x1, x2, xn = b1, b2, abn
	}
}
