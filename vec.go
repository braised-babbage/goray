package main

import "math"

type Vec3 [3]float64

func (v *Vec3) X() float64 { return v[0] }
func (v *Vec3) Y() float64 { return v[1] }
func (v *Vec3) Z() float64 { return v[2] }

func (v *Vec3) Neg() Vec3 { return Vec3{-v[0], -v[1], -v[2]} }
func (v *Vec3) Increment(w Vec3) *Vec3 {
	v[0] += w[0]
	v[1] += w[1]
	v[2] += w[2]
	return v
}
func (v *Vec3) Decrement(w Vec3) *Vec3 {
	v[0] -= w[0]
	v[1] -= w[1]
	v[2] -= w[2]
	return v
}
func (v *Vec3) Scale(t float64) *Vec3 {
	v[0] *= t
	v[1] *= t
	v[2] *= t
	return v
}
func (v *Vec3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}
func (v *Vec3) LengthSquared() float64 {
	return v[0]*v[0] + v[1]*v[1] + v[2]*v[2]
}
func (v *Vec3) Normalized() Vec3 {
	l := v.Length()
	return Vec3{v[0] / l, v[1] / l, v[2] / l}
}

func (v *Vec3) Add(w *Vec3) *Vec3 {
	return &Vec3{v[0] + w[0], v[1] + w[1], v[2] + w[2]}
}
func (v *Vec3) Sub(w *Vec3) *Vec3 {
	return &Vec3{v[0] - w[0], v[1] - w[1], v[2] - w[2]}
}
func (v *Vec3) Mul(t float64) *Vec3 {
	return &Vec3{v[0] * t, v[1] * t, v[2] * t}
}

type Point3 = Vec3
type RGB = Vec3

func Interp(u Vec3, v Vec3, t float64) Vec3 {
	s := 1.0 - t
	return Vec3{s*u[0] + t*v[0], s*u[1] + t*v[1], s*u[2] + t*v[2]}
}

func Axpy(t float64, x Vec3, y Vec3) Vec3 {
	return Vec3{t*x[0] + y[0], t*x[1] + y[1], t*x[2] + y[2]}
}

func UVCoord(orig Vec3, udir Vec3, vdir Vec3, u float64, v float64) Vec3 {
	return Vec3{
		orig[0] + udir[0]*u + vdir[0]*v,
		orig[1] + udir[1]*u + vdir[1]*v,
		orig[2] + udir[2]*u + vdir[2]*v,
	}
}
