package main

import (
	"image/png"
	"os"

	"github.com/go-micah/imaging"
)

func main() {

	img := imaging.Steps(1033, 512)
	f, _ := os.Create("./output/steps.png")
	png.Encode(f, img)

	img = imaging.Dist(1033, 512)
	f, _ = os.Create("./output/dist.png")
	png.Encode(f, img)

}
