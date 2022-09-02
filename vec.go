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

type Point3 = Vec3
type RGB = Vec3
