package plotC

import (
	"image"
	col "image/color"
	"image/png"
	"log"
	mt "math"
	cm "math/cmplx"
	"os"
)

var S int = 1024
var L float64 = 10
var X0 float64 = 0
var Y0 float64 = 0
var O byte = 'c'

func Poly_(C []float64) func(complex128) complex128 {
	return func(x complex128) complex128 {
		var r, X complex128 = 0, 1
		for i := 0; i < len(C); i++ {
			r += mulC(X, C[i])
			X *= x
		}
		return r
	}
}

func atan(x, y float64) float64 {
	return mt.Mod(mt.Atan2(y, x)+2*mt.Pi, 2*mt.Pi)
}

func top(x, y float64) float64 {
	if x <= y {
		return x
	}
	return y
}

//hollow knight
func Color(x, y float64, f func(complex128) complex128) []uint8 { //retorna un color segun el numero complejo que retorna f
	C := f(complex(x, y))
	c := [2]float64{0, 0}
	c[0] = cm.Abs(C)
	c[1] = atan(real(C), imag(C))
	var r, g, b float64 = 0, 0, 0
	var m float64 = (510 * c[0] / (c[0] + 1))
	if m == 0 {
		return []uint8{0, 0, 0}
	}
	if c[1] < 2*mt.Pi/3 {
		g = (m * 3 * c[1] / (2 * mt.Pi))
		r = m - g
	} else {
		if c[1] < 4*mt.Pi/3 {
			b = m * 3 * (c[1] - 2*mt.Pi/3) / (2 * mt.Pi)
			g = m - b
		} else {
			r = m * 3 * (c[1] - 4*mt.Pi/3) / (2 * mt.Pi)
			b = m - r
		}
	}
	return []uint8{uint8(top(r, 255)), uint8(top(g, 255)), uint8(top(b, 255))}
}

func Gray(x, y float64, f func(complex128) complex128) []uint8 {
	C := f(complex(x, y))
	c := [2]float64{0, 0}
	c[0] = cm.Abs(C)
	c[1] = atan(real(C), imag(C))
	var m float64 = (255 * c[0] / (c[0] + 1))
	return []uint8{uint8(m), uint8(m), uint8(m)}
}

func Plot(f_ func(complex128) complex128) *image.NRGBA {
	Sf := float64(S)
	x0 := X0
	y0 := Y0
	img := image.NewNRGBA(image.Rect(0, 0, S, S))
	var C []uint8
	var x, y float64 = 0, 0
	p := f_
	for i := 0; i < S; i++ {
		x = (float64(i)-Sf/float64(2))*L/Sf + x0
		for j := 0; j < S; j++ {
			y = (float64(j)-Sf/2)*L/Sf + y0
			switch O {
			case 'c':
				C = Color(x, y, p)
			case 'g':
				C = Gray(x, y, p)
			}
			img.Set(i, j, col.NRGBA{R: C[0], G: C[1], B: C[2], A: 255})
		}
	}
	return img
}

func Save(img *image.NRGBA, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	if err := png.Encode(f, image.Image(img)); err != nil {
		f.Close()
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

//______________________________________________________________________
func cuad(C []float64) []float64 { //C->C2
	L := len(C)
	r := []float64{}
	k := float64(0)
	for i := 0; i < 2*L-1; i++ {
		k = float64(0)
		if i < L {
			for j := 0; j < i+1; j++ {
				k += C[j] * C[i-j]
			}
			r = append(r, k)
		} else {
			for j := i - L + 1; j < L; j++ {
				k += C[j] * C[i-j]
			}
			r = append(r, k)
		}
	}
	return r
}
