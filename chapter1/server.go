package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lissajous", picture)
	http.HandleFunc("/svg", svg)
	log.Fatal(http.ListenAndServe("localhost:8088", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.Path, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s] = %s\n", k, v)
	}
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Counter = %d\n", count)
	mu.Unlock()
}

func picture(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
		return
	}
	fmt.Println(cycles)
	lissajous(w, cycles)
}

func svg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	draw(w)
}

var palette = []color.Color{color.White, color.RGBA{0xff, 0x00, 0xff, 0xff}}

const (
	whiteIndex = 0
	greenIndex = 1
)

func lissajous(out io.Writer, cycles int) {
	const (
		// cycles  = 5     // 周期数
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
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
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

// -------------------------------------------------------------------- //

const (
	width, height = 600, 320            // 画布宽高像素数
	cells         = 100                 // 方格数
	xyrange       = 30.0                // xy的范围 (-xyrange, xyrange)
	xyscale       = width / 2 / xyrange // 单位长度内x或y的像素数
	zscale        = height * .4         // 单位长度内z的像素数
	angle         = math.Pi / 6         // x, y轴的角度
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func draw(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
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
