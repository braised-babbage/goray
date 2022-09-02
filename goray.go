package main

import (
	"fmt"
	"os"
)

func RayColor(r Ray) RGB {
	unit_dir := r.dir.Normalized()
	t := 0.5 * (unit_dir.Y() + 1.0)
	return Interp(RGB{1.0, 1.0, 1.0}, RGB{0.5, 0.7, 1.0}, t)
}

func main() {

	// image
	const aspect_ratio = 16.0 / 9.0
	const image_width = 256
	const image_height = image_width / aspect_ratio

	// camera
	viewport_height := 2.0
	viewport_width := aspect_ratio * viewport_height
	focal_length := 1.0

	origin := Point3{0, 0, 0}
	horizontal := Vec3{viewport_width, 0.0, 0.0}
	vertical := Vec3{0.0, viewport_height, 0.0}
	lower_left := UVCoord(origin, horizontal, vertical, -0.5, -0.5)
	lower_left.Decrement(Vec3{0, 0, focal_length})

	// render
	fmt.Printf("P3\n%v %v\n255\n", image_width, image_height)

	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %v", j)

		for i := 0; i < image_width; i++ {
			u := float64(i) / float64(image_width-1)
			v := float64(j) / float64(image_height-1)
			target := UVCoord(lower_left, horizontal, vertical, u, v)
			r := Ray{origin, *target.Decrement(origin)}
			pixel_color := RayColor(r)
			WriteRGB(os.Stdout, pixel_color)
		}
	}
}
