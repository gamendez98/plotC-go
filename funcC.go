package plotC

import (
	mt "math"
	cm "math/cmplx"
)

/*
var P = []float64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47,
	53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109,
	113, 127, 131, 137, 139, 149, 151, 157, 163, 167,
	173, 179, 181, 191, 193, 197, 199, 211, 223, 227,
	229, 233, 239, 241, 251, 257, 263, 269, 271, 277,
	281, 283, 293, 307, 311, 313, 317, 331, 337, 347,
	349, 353, 359, 367, 373, 379, 383, 389, 397, 401,
	409, 419, 421, 431, 433, 439, 443, 449, 457, 461,
	463, 467, 479, 487, 491, 499, 503, 509, 521, 523,
	541, 547, 557, 563, 569, 571, 577, 587, 593, 599}
*/
func mulC(c complex128, z float64) complex128 {
	return complex(real(c)*z, imag(c)*z)
}

func exp(x float64, y complex128) complex128 { //returns 1/x^y
	return cm.Exp(mulC(y, mt.Log(x)))
}

func Gamma(x complex128) complex128 {
	uc, r := complex128(1), complex128(1)
	N, uf := float64(50), float64(1)
	if x == 0 {
		return 0
	}
	for i := uf; i < N; i++ {
		r = r * exp(uf+uf/i, x) / (uc + mulC(x, uf/i))
	}
	return r / x
}

func Mandel_(n int) func(complex128) complex128 {
	return func(x complex128) complex128 {
		r := complex128(0)
		for i := 0; i < n; i++ {
			r = r*r + x
		}
		return r
	}
}

func Julia_(n int, c complex128) func(complex128) complex128 {
	return func(x complex128) complex128 {
		r := x
		for i := 0; i < n; i++ {
			r = r*r + c
		}
		return r
	}
}

func Spin(x complex128) complex128 {
	M := mt.Log(cm.Abs(x))
	return x * complex(mt.Sin(M), mt.Cos(M))
}
