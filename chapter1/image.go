package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.RGBA{0xff, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0
	greenIndex = 1
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 周期数
		res     = 0.001 // 角分辨率
		size    = 100   // 图片范围为[-size, +size]
		nframes = 64    // 动画帧数
		delay   = 8     // 帧之间的间隔，以10ms计
	)

	freq := rand.Float64() * 3.0 // 相对于y的频率
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //相位差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
