package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func runSin()  {

	const size = 300

	pic := image.NewGray(image.Rect(0,0,size,size))

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			pic.SetGray(i,j,color.Gray{255})
		}
	}

	for x := 0; x < size; x++ {
		s := float64(x) * 2 * math.Pi / size
		//y := size / 2 - math.Sin(s) * size / 2
		y := size / 2 - math.Sin(s) * size / 2
		pic.SetGray(x,int(y),color.Gray{0})
	}

	file,err := os.Create("sin.png")

	if err != nil {
		log.Fatal(err)
	}

	png.Encode(file,pic)

	file.Close()

}
