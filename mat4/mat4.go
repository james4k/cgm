package mat4

import (
	"j4k.co/cgm"
	"j4k.co/cgm/vec3"
)

func Identity() [16]float32 {
	return [16]float32{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

func Mul(a, b [16]float32) [16]float32 {
	return [16]float32{
		b[0]*a[0] + b[1]*a[4] + b[2]*a[8] + b[3]*a[12],
		b[0]*a[1] + b[1]*a[5] + b[2]*a[9] + b[3]*a[13],
		b[0]*a[2] + b[1]*a[6] + b[2]*a[10] + b[3]*a[14],
		b[0]*a[3] + b[1]*a[7] + b[2]*a[11] + b[3]*a[15],
		b[4]*a[0] + b[5]*a[4] + b[6]*a[8] + b[7]*a[12],
		b[4]*a[1] + b[5]*a[5] + b[6]*a[9] + b[7]*a[13],
		b[4]*a[2] + b[5]*a[6] + b[6]*a[10] + b[7]*a[14],
		b[4]*a[3] + b[5]*a[7] + b[6]*a[11] + b[7]*a[15],
		b[8]*a[0] + b[9]*a[4] + b[10]*a[8] + b[11]*a[12],
		b[8]*a[1] + b[9]*a[5] + b[10]*a[9] + b[11]*a[13],
		b[8]*a[2] + b[9]*a[6] + b[10]*a[10] + b[11]*a[14],
		b[8]*a[3] + b[9]*a[7] + b[10]*a[11] + b[11]*a[15],
		b[12]*a[0] + b[13]*a[4] + b[14]*a[8] + b[15]*a[12],
		b[12]*a[1] + b[13]*a[5] + b[14]*a[9] + b[15]*a[13],
		b[12]*a[2] + b[13]*a[6] + b[14]*a[10] + b[15]*a[14],
		b[12]*a[3] + b[13]*a[7] + b[14]*a[11] + b[15]*a[15],
	}
}

func Mul4(m [16]float32, v [4]float32) [4]float32 {
	return [4]float32{
		v[0]*m[0] + v[1]*m[4] + v[2]*m[8] + v[3]*m[12],
		v[0]*m[1] + v[1]*m[5] + v[2]*m[9] + v[3]*m[13],
		v[0]*m[2] + v[1]*m[6] + v[2]*m[10] + v[3]*m[14],
		v[0]*m[3] + v[1]*m[7] + v[2]*m[11] + v[3]*m[15],
	}
}

func Mul3(m [16]float32, v [3]float32) [3]float32 {
	return [3]float32{
		v[0]*m[0] + v[1]*m[4] + v[2]*m[8],
		v[0]*m[1] + v[1]*m[5] + v[2]*m[9],
		v[0]*m[2] + v[1]*m[6] + v[2]*m[10],
	}
}

// Inv returns the inverse of m. Panics if Determinant of m is zero.
func Inv(m [16]float32) [16]float32 {
	invd := 1 / Det(m)
	return [16]float32{
		(-m[7]*m[10]*m[13] + m[6]*m[11]*m[13] + m[7]*m[9]*m[14] -
			m[5]*m[11]*m[14] - m[6]*m[9]*m[15] + m[5]*m[10]*m[15]) * invd,
		(m[3]*m[10]*m[13] - m[2]*m[11]*m[13] - m[3]*m[9]*m[14] +
			m[1]*m[11]*m[14] + m[2]*m[9]*m[15] - m[1]*m[10]*m[15]) * invd,
		(-m[3]*m[6]*m[13] + m[2]*m[7]*m[13] + m[3]*m[5]*m[14] -
			m[1]*m[7]*m[14] - m[2]*m[5]*m[15] + m[1]*m[6]*m[15]) * invd,
		(m[3]*m[6]*m[9] - m[2]*m[7]*m[9] - m[3]*m[5]*m[10] +
			m[1]*m[7]*m[10] + m[2]*m[5]*m[11] - m[1]*m[6]*m[11]) * invd,
		(m[7]*m[10]*m[12] - m[6]*m[11]*m[12] - m[7]*m[8]*m[14] +
			m[4]*m[11]*m[14] + m[6]*m[8]*m[15] - m[4]*m[10]*m[15]) * invd,
		(-m[3]*m[10]*m[12] + m[2]*m[11]*m[12] + m[3]*m[8]*m[14] -
			m[0]*m[11]*m[14] - m[2]*m[8]*m[15] + m[0]*m[10]*m[15]) * invd,
		(m[3]*m[6]*m[12] - m[2]*m[7]*m[12] - m[3]*m[4]*m[14] +
			m[0]*m[7]*m[14] + m[2]*m[4]*m[15] - m[0]*m[6]*m[15]) * invd,
		(-m[3]*m[6]*m[8] + m[2]*m[7]*m[8] + m[3]*m[4]*m[10] -
			m[0]*m[7]*m[10] - m[2]*m[4]*m[11] + m[0]*m[6]*m[11]) * invd,
		(-m[7]*m[9]*m[12] + m[5]*m[11]*m[12] + m[7]*m[8]*m[13] -
			m[4]*m[11]*m[13] - m[5]*m[8]*m[15] + m[4]*m[9]*m[15]) * invd,
		(m[3]*m[9]*m[12] - m[1]*m[11]*m[12] - m[3]*m[8]*m[13] +
			m[0]*m[11]*m[13] + m[1]*m[8]*m[15] - m[0]*m[9]*m[15]) * invd,
		(-m[3]*m[5]*m[12] + m[1]*m[7]*m[12] + m[3]*m[4]*m[13] -
			m[0]*m[7]*m[13] - m[1]*m[4]*m[15] + m[0]*m[5]*m[15]) * invd,
		(m[3]*m[5]*m[8] - m[1]*m[7]*m[8] - m[3]*m[4]*m[9] +
			m[0]*m[7]*m[9] + m[1]*m[4]*m[11] - m[0]*m[5]*m[11]) * invd,
		(m[6]*m[9]*m[12] - m[5]*m[10]*m[12] - m[6]*m[8]*m[13] +
			m[4]*m[10]*m[13] + m[5]*m[8]*m[14] - m[4]*m[9]*m[14]) * invd,
		(-m[2]*m[9]*m[12] + m[1]*m[10]*m[12] + m[2]*m[8]*m[13] -
			m[0]*m[10]*m[13] - m[1]*m[8]*m[14] + m[0]*m[9]*m[14]) * invd,
		(m[2]*m[5]*m[12] - m[1]*m[6]*m[12] - m[2]*m[4]*m[13] +
			m[0]*m[6]*m[13] + m[1]*m[4]*m[14] - m[0]*m[5]*m[14]) * invd,
		(-m[2]*m[5]*m[8] + m[1]*m[6]*m[8] + m[2]*m[4]*m[9] -
			m[0]*m[6]*m[9] - m[1]*m[4]*m[10] + m[0]*m[5]*m[10]) * invd,
	}
}

// Det returns the determinant of m.
func Det(m [16]float32) float32 {
	return m[0]*m[5]*m[10]*m[15] - m[0]*m[5]*m[11]*m[14] -
		m[0]*m[6]*m[9]*m[15] + m[0]*m[6]*m[11]*m[13] +
		m[0]*m[7]*m[9]*m[14] - m[0]*m[7]*m[10]*m[13] -
		m[1]*m[4]*m[10]*m[15] + m[1]*m[4]*m[11]*m[14] +
		m[1]*m[6]*m[8]*m[15] - m[1]*m[6]*m[11]*m[12] -
		m[1]*m[7]*m[8]*m[14] + m[1]*m[7]*m[10]*m[12] +
		m[2]*m[4]*m[9]*m[15] - m[2]*m[4]*m[11]*m[13] -
		m[2]*m[5]*m[8]*m[15] + m[2]*m[5]*m[11]*m[12] +
		m[2]*m[7]*m[8]*m[13] - m[2]*m[7]*m[9]*m[12] -
		m[3]*m[4]*m[9]*m[14] + m[3]*m[4]*m[10]*m[13] +
		m[3]*m[5]*m[8]*m[14] - m[3]*m[5]*m[10]*m[12] -
		m[3]*m[6]*m[8]*m[13] + m[3]*m[6]*m[9]*m[12]
}

func LookAtLH(eye, at, up [3]float32) [16]float32 {
	z := vec3.Normal(vec3.Sub(at, eye))
	x := vec3.Normal(vec3.Cross(up, z))
	y := vec3.Cross(z, x)
	return [16]float32{
		x[0], y[0], z[0], 0,
		x[1], y[1], z[1], 0,
		x[2], y[2], z[2], 0,
		-vec3.Dot(x, eye),
		-vec3.Dot(y, eye),
		-vec3.Dot(z, eye),
		1.0,
	}
}

func LookAt(eye, at, up [3]float32) [16]float32 {
	z := vec3.Normal(vec3.Sub(at, eye))
	x := vec3.Normal(vec3.Cross(z, up))
	y := vec3.Cross(x, z)
	return [16]float32{
		x[0], y[0], -z[0], 0,
		x[1], y[1], -z[1], 0,
		x[2], y[2], -z[2], 0,
		-vec3.Dot(x, eye),
		-vec3.Dot(y, eye),
		vec3.Dot(z, eye),
		1.0,
	}
}

func PerspectiveLH(fovY cgm.Radians, aspectRatio, nearZ, farZ float32) [16]float32 {
	f := 1.0 / cgm.Tan(fovY*0.5)
	d := farZ - nearZ
	return [16]float32{
		f / aspectRatio, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (farZ + nearZ) / d, 1,
		0, 0, -2 * nearZ * farZ / d, 0,
	}
}

// Perspective computes a right-handed projection matrix, mapping Z to
// [-1, 1], suitable for OpenGL-like pipelines.
func Perspective(fovY cgm.Radians, aspectRatio, nearZ, farZ float32) [16]float32 {
	f := 1.0 / cgm.Tan(fovY*0.5)
	d := nearZ - farZ
	return [16]float32{
		f / aspectRatio, 0, 0, 0,
		0, f, 0, 0,
		0, 0, (farZ + nearZ) / d, -1,
		0, 0, 2 * nearZ * farZ / d, 0,
	}
}

// PerspectiveD3D computes a right-handed projection matrix, mapping Z
// to [0, 1], suitable for Direct3D-like pipelines.
func PerspectiveD3D(fovY cgm.Radians, aspectRatio, nearZ, farZ float32) [16]float32 {
	f := 1.0 / cgm.Tan(fovY*0.5)
	d := nearZ - farZ
	return [16]float32{
		f / aspectRatio, 0, 0, 0,
		0, f, 0, 0,
		0, 0, farZ / d, -1,
		0, 0, nearZ * farZ / d, 0,
	}
}

func OrthoLH(left, right, bottom, top, near, far float32) [16]float32 {
	return [16]float32{
		0:  2 / (right - left),
		5:  2 / (top - bottom),
		10: 1 / (far - near),
		12: (left + right) / (left - right),
		(top + bottom) / (bottom - top),
		near / (near - far),
		1,
	}
}

func RotateXYZ(x, y, z cgm.Radians) [16]float32 {
	var (
		sx = cgm.Sin(x)
		sy = cgm.Sin(y)
		sz = cgm.Sin(z)
		cx = cgm.Cos(x)
		cy = cgm.Cos(y)
		cz = cgm.Cos(z)
	)
	return [16]float32{
		cy * cz, -cy * sz, sy, 0,
		cz*sx*sy + cx*sz, cx*cz - sx*sy*sz, -cy * sx, 0,
		-cx*cz*sy + sx*sz, cz*sx + cx*sy*sz, cx * cy, 0,
		0, 0, 0, 1,
	}
}
