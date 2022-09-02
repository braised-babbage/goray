package main

type Ray struct {
	origin Point3
	dir    Vec3
}

func (r Ray) At(t float64) Point3 {
	return Axpy(t, r.dir, r.origin)
}
