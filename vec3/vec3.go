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

func Mul(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[0] * b[0],
		a[1] * b[1],
		a[2] * b[2],
	}
}

// Mul1 multiples vector a with scalar b.
func Mul1(a [3]float32, b float32) [3]float32 {
	return [3]float32{
		a[0] * b,
		a[1] * b,
		a[2] * b,
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
	return Mul1(a, 1/Length(a))
}
