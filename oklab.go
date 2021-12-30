// Package oklab is a Go port of Björn Ottosson's C++ code
// that converts linear sRGB to Oklab color space and vice versa.
package oklab

import "math"

type RGB struct {
	R, G, B float32
}

type Lab struct {
	L, A, B float32
}

func LinearRgbToOklab(c RGB) Lab {
	r, g, b := float64(c.R), float64(c.G), float64(c.B)
	lng := 0.4122214708*r + 0.5363325363*g + 0.0514459929*b
	mdl := 0.2119034982*r + 0.6806995451*g + 0.1073969566*b
	sht := 0.0883024619*r + 0.2817188376*g + 0.6299787005*b

	lRoot := math.Cbrt(lng)
	mRoot := math.Cbrt(mdl)
	sRoot := math.Cbrt(sht)

	lt := 0.2104542553*lRoot + 0.7936177850*mRoot - 0.0040720468*sRoot
	ca := 1.9779984951*lRoot - 2.4285922050*mRoot + 0.4505937099*sRoot
	cb := 0.0259040371*lRoot + 0.7827717662*mRoot - 0.8086757660*sRoot
	return Lab{float32(lt), float32(ca), float32(cb)}
}

func OklabToLinearRgb(c Lab) RGB {
	l, a, b := float64(c.L), float64(c.A), float64(c.B)
	lRoot := l + 0.3963377774*a + 0.2158037573*b
	mRoot := l - 0.1055613458*a - 0.0638541728*b
	sRoot := l - 0.0894841775*a - 1.2914855480*b

	long := lRoot * lRoot * lRoot
	middle := mRoot * mRoot * mRoot
	short := sRoot * sRoot * sRoot

	rd := +4.0767416621*long - 3.3077115913*middle + 0.2309699292*short
	gr := -1.2684380046*long + 2.6097574011*middle - 0.3413193965*short
	bl := -0.0041960863*long - 0.7034186147*middle + 1.7076147010*short
	return RGB{float32(rd), float32(gr), float32(bl)}
}
