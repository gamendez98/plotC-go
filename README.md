# plotC-go
a package for ploting complex functions with a color representation

# plotC.go:
main functionalities

# functions:

func Poly_(C []float64) func(complex128) complex128:
creates a polynomial function from its coheficients

func Color(x, y float64, f func(complex128) complex128) []uint8 :
gives a color representation of a complex number

func Gray(x, y float64, f func(complex128) complex128) []uint8:
returns a gray scale representation of a complex number's magnitud

func Plot(f_ func(complex128) complex128) *image.NRGBA:
gives an image that represents the values of a complex function over a rectangular section of the complex plain

func Save(img *image.NRGBA, file string) :
saves the image in a png format
file must end in png(this is not cheked by the function)

# options:

var S int = 1024:
the definition of the image returned by Plot

var L float64 = 10:
sets the size of the plotet section (taken from the center to the sides)

var X0 float64 = 0
var Y0 float64 = 0:
the center point of the image in the complex plain

var O byte = 'c':
'c':for ploting with Color
'g':for ploting with Gray

# funcC.go:
some complex functions

# functions:

func Gamma(x complex128) complex128:
aproximation of the gamma function

func Mandel_(n int) func(complex128) complex128 :
returns a function that is the nth iteration of z_n+1=z²+x where x is the input of the returned function and z_0=0

func Julia_(n int, c complex128) func(complex128) complex128:
returns a function that is the nth iteration of z_n+1=z²+c where z_0 is the input of the returned function

func Spin(x complex128) complex128 :
this function twist the complex plain

