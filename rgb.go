package main

import (
	"fmt"
	"io"
)

const cscale = 255.999

func WriteRGB(w io.Writer, c RGB) {
	fmt.Fprintf(w, "%v %v %v\n", int(c[0]*cscale), int(c[1]*cscale), int(c[2]*cscale))
}
