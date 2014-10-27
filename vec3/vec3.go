package vec3

import "j4k.co/cgm"

func Add(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[0] + b[0],
		a[1] + b[1],
		a[2] + b[2],
	}
}

func Sub(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[0] - b[0],
		a[1] - b[1],
		a[2] - b[2],
	}
}

// Mul multiples scalar k by vector v.
func Mul(k float32, v [3]float32) [3]float32 {
	return [3]float32{
		k * v[0],
		k * v[1],
		k * v[2],
	}
}

func Dot(a, b [3]float32) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func Cross(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func Length(a [3]float32) float32 {
	return cgm.Sqrt(Dot(a, a))
}

func Normal(a [3]float32) [3]float32 {
	return Mul(1/Length(a), a)
}
