package main

import (
	"fmt"
	"os"
)

func main() {
	const image_width = 256
	const image_height = 256

	fmt.Printf("P3\n%v %v\n255\n", image_width, image_height)

	for j := image_height - 1; j >= 0; j-- {
		fmt.Fprintf(os.Stderr, "\rScanlines remaining: %v", j)

		for i := 0; i < image_width; i++ {
			r := float64(i) / float64(image_width-1)
			g := float64(j) / float64(image_height-1)
			b := 0.25

			ir := int(255.999 * r)
			ig := int(255.999 * g)
			ib := int(255.999 * b)

			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}
