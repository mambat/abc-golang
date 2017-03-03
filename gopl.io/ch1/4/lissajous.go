// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
    "image/color"
    "math/rand"
    "time"
    "os"
    "io"
    "image/gif"
    "image"
    "math"
)

// 复合类型: slice切片
//var palette = []color.Color{color.White, color.Black}
var palette = []color.Color{color.White, color.Black, color.RGBA{0, 128, 0, 0xff}, color.RGBA{128, 0, 0, 0xff}}

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    // Thanks to Randall McPherson for pointing out the omission.
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    // 目前常量声明的值必须是一个数字值、字符串或者一个固定的boolean值
    const (
        cycles  = 5     // number of complete x oscillator revolutions
        res     = 0.001 // angular resolution
        size    = 100   // image canvas covers [-size..+size]
        nframes = 128    // number of animation frames
        delay   = 8     // delay between frames in 10ms units
    )

    freq := rand.Float64() * 3.0 // relative frequency of y oscillator

    // 复合类型: struct结构体
    anim := gif.GIF{LoopCount: nframes}

    phase := 0.0 // phase difference

    for i, c := 0, 1; i < nframes; i++ {
        rect := image.Rect(0, 0, 2*size+1, 2*size+1)
        img := image.NewPaletted(rect, palette)
        for t := 0.0; t < cycles*2*math.Pi; t += res {
            x := math.Sin(t)
            y := math.Sin(t*freq + phase)
            img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(c))
        }

        c++ // change color
        if c > len(palette) {
            c = 1
        }

        phase += 0.1
        anim.Delay = append(anim.Delay, delay)
        anim.Image = append(anim.Image, img)
    }
    gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
