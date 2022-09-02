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
			pixel_color := RGB{
				float64(i) / float64(image_height-1),
				float64(j) / float64(image_width-1),
				0.25,
			}
			WriteRGB(os.Stdout, pixel_color)
		}
	}
}
