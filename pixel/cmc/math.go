package main

import (
	"math"
)

const TAU = 2 * math.Pi

func Clamp1(a float32) float32 {
	if a > 1 {
		return 1
	} else if a < 0 {
		return 0
	}
	return a
}
func Sincos(a float32) (float32, float32) {
	sn, cs := math.Sincos(float64(a))
	return float32(sn), float32(cs)
}
func Sin(a float32) float32  { return float32(math.Sin(float64(a))) }
func Cos(a float32) float32  { return float32(math.Cos(float64(a))) }
func Sqrt(a float32) float32 { return float32(math.Sqrt(float64(a))) }

func Lerp(a, b, p float32) float32 { return a + (b-a)*p }

func Log(v float32) float32 { return float32(math.Log(float64(v))) }
func Exp(v float32) float32 { return float32(math.Exp(float64(v))) }

func Atan2(y, x float32) float32 { return float32(math.Atan2(float64(y), float64(x))) }

func LogX16(v uint16) float32 { return logx16[v] }

var logx16 [0xFFFF]float32

func init() {
	for i := range logx16 {
		logx16[i] = float32(math.Log(float64(i) + 1))
	}
}
