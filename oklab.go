package goklab

import "math"

type RGB struct {
	R, G, B float64
}

type Lab struct {
	L, A, B float64
}

func LinearRgbToOklab(c RGB) Lab {
	lng := 0.4122214708*c.R + 0.5363325363*c.G + 0.0514459929*c.B
	mdl := 0.2119034982*c.R + 0.6806995451*c.G + 0.1073969566*c.B
	sht := 0.0883024619*c.R + 0.2817188376*c.G + 0.6299787005*c.B

	lRoot := math.Cbrt(lng)
	mRoot := math.Cbrt(mdl)
	sRoot := math.Cbrt(sht)

	lt := 0.2104542553*lRoot + 0.7936177850*mRoot - 0.0040720468*sRoot
	ca := 1.9779984951*lRoot - 2.4285922050*mRoot + 0.4505937099*sRoot
	cb := 0.0259040371*lRoot + 0.7827717662*mRoot - 0.8086757660*sRoot
	return Lab{lt, ca, cb}
}

func OklabToLinearRgb(c Lab) RGB {
	lRoot := c.L + 0.3963377774*c.A + 0.2158037573*c.B
	mRoot := c.L - 0.1055613458*c.A - 0.0638541728*c.B
	sRoot := c.L - 0.0894841775*c.A - 1.2914855480*c.B

	long := lRoot * lRoot * lRoot
	middle := mRoot * mRoot * mRoot
	short := sRoot * sRoot * sRoot

	rd := +4.0767416621*long - 3.3077115913*middle + 0.2309699292*short
	gr := -1.2684380046*long + 2.6097574011*middle - 0.3413193965*short
	bl := -0.0041960863*long - 0.7034186147*middle + 1.7076147010*short
	return RGB{rd, gr, bl}
}
