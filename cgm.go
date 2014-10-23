// Package cgm is a math library with a flavor for computer graphics.
// Incomplete and likely to change drastically.
package cgm

import "math"

func Sqrt(x float32) float32 {
	return float32(math.Sqrt(float64(x)))
}

type Degrees float32

func (d Degrees) ToRadians() Radians {
	return Radians(d * math.Pi / 180)
}

type Radians float32

func (r Radians) ToDegrees() Degrees {
	return Degrees(r * 180 / math.Pi)
}

func Sin(r Radians) float32 {
	return float32(math.Sin(float64(r)))
}

func Cos(r Radians) float32 {
	return float32(math.Cos(float64(r)))
}

func Tan(r Radians) float32 {
	return float32(math.Tan(float64(r)))
}
