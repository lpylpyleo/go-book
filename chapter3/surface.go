package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // 画布宽高像素数
	cells         = 100                 // 方格数
	xyrange       = 30.0                // xy的范围 (-xyrange, xyrange)
	xyscale       = width / 2 / xyrange // 单位长度内x或y的像素数
	zscale        = height * .4         // 单位长度内z的像素数
	angle         = math.Pi / 6         // x, y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// 算出格子(i, j)的角点(x, y)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// 3d投影
	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
