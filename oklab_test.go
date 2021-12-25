package oklab

import (
	"math"
	"testing"
)

const eps = 0.0000011

type testCase struct {
	rgb RGB
	lab Lab
}

func getData() []testCase {
	return []testCase{
		{
			RGB{1.000000e+00, 0.000000e+00, 0.000000e+00},
			Lab{6.279554e-01, 2.248631e-01, 1.258463e-01},
		},
		{
			RGB{0.000000e+00, 1.000000e+00, 0.000000e+00},
			Lab{8.664396e-01, -2.338874e-01, 1.794985e-01},
		},
		{
			RGB{0.000000e+00, 0.000000e+00, 1.000000e+00},
			Lab{4.520137e-01, -3.245696e-02, -3.115281e-01},
		},
		{
			RGB{1.000000e+00, 1.000000e+00, 1.000000e+00},
			Lab{1.000000e+00, 0.000000e+00, 5.960464e-08},
		},
		{
			RGB{0.000000e+00, 0.000000e+00, 0.000000e+00},
			Lab{0.000000e+00, 0.000000e+00, 0.000000e+00},
		},
		{
			RGB{2.462013e-01, 6.583748e-01, 2.232280e-01},
			Lab{8.002464e-01, -1.078115e-01, 8.255047e-02},
		},
		{
			RGB{1.000000e+00, 3.467041e-01, 3.954624e-02},
			Lab{7.842250e-01, 7.268986e-02, 1.413031e-01},
		},
		{
			RGB{5.520114e-01, 5.149177e-01, 3.813260e-01},
			Lab{7.999293e-01, -3.195465e-03, 3.335571e-02},
		},
		{
			RGB{1.000000e+00, 5.028865e-01, 1.000000e+00},
			Lab{8.774487e-01, 9.605274e-02, -6.356859e-02},
		},
		{
			RGB{1.000000e+00, 6.048833e-03, 2.050787e-01},
			Lab{6.477489e-01, 2.548900e-01, 1.702711e-02},
		},
		{
			RGB{1.000000e+00, 0.000000e+00, 7.036010e-02},
			Lab{6.337076e-01, 2.412873e-01, 7.909188e-02},
		},
	}
}

func TestItIsReversible(t *testing.T) {
	data := getData()
	for i := 0; i < len(data); i++ {
		input := data[i].rgb
		lab := LinearRgbToOklab(input)
		result := OklabToLinearRgb(lab)

		if math.Abs(float64(result.R-input.R)) >= eps {
			t.Fatalf("#%d R: %e != %e", i, result.R, input.R)
		}

		if math.Abs(float64(result.G-input.G)) >= eps {
			t.Fatalf("#%d G: %e != %e", i, result.G, input.G)
		}

		if math.Abs(float64(result.B-input.B)) >= eps {
			t.Fatalf("#%d B: %e != %e", i, result.B, input.B)
		}
	}
}

func toSRGB(linear float64) float64 {
	const a = 0.055
	if linear <= 0.0031308 {
		return linear * 12.92
	} else {
		return math.Pow(float64(linear), 1.0/2.4)*(1+a) - a
	}
}

func toLinear(sRGB float64) float64 {
	const a = 0.055
	if sRGB <= 0.04045 {
		return sRGB / 12.92
	} else {
		return math.Pow((sRGB+a)/(1+a), 2.4)
	}
}

func TestIntegerSrgbReversible(t *testing.T) {
	minR, maxR := 128.0, 128.0
	minG, maxG := 128.0, 128.0
	minB, maxB := 128.0, 128.0

	for i := 0x0; i <= 0xFF_FF_FF; i++ {
		r := i >> 16
		g := (i % 0x1_00_00) >> 8
		b := i % 0x1_00

		minR, maxR = math.Min(minR, float64(r)), math.Max(maxR, float64(r))
		minG, maxG = math.Min(minG, float64(g)), math.Max(maxG, float64(g))
		minB, maxB = math.Min(minB, float64(b)), math.Max(maxB, float64(b))

		input := RGB{
			float32(toLinear(float64(r) / 255.0)),
			float32(toLinear(float64(g) / 255.0)),
			float32(toLinear(float64(b) / 255.0))}

		output := OklabToLinearRgb(LinearRgbToOklab(input))

		rr := int(math.Round(255.0 * toSRGB(float64(output.R))))
		rg := int(math.Round(255.0 * toSRGB(float64(output.G))))
		rb := int(math.Round(255.0 * toSRGB(float64(output.B))))

		if (rr != r) || (rg != g) || (rb != b) {
			t.Fatalf("Input: %d %d %d. Output: %d %d %d.",
				r, g, b, rr, rg, rb)
		}
	}

	if (math.Abs(minR) >= eps) ||
		(math.Abs(minG) >= eps) ||
		(math.Abs(minB) >= eps) ||
		(math.Abs(maxR-255.0) >= eps) ||
		(math.Abs(maxG-255.0) >= eps) ||
		(math.Abs(maxB-255.0) >= eps) {
		t.Fatalf("This test is broken. MinRGB: %f %f %f. MaxRGB: %f %f %f.",
			minR, minG, minB, maxR, maxG, maxB)
	}
}

func TestRgbToLab(t *testing.T) {
	data := getData()
	for i := 0; i < len(data); i++ {
		input := data[i].rgb
		expected := data[i].lab
		result := LinearRgbToOklab(input)

		if math.Abs(float64(result.L-expected.L)) >= eps {
			t.Fatalf("#%d L: %e != %e", i, result.L, expected.L)
		}

		if math.Abs(float64(result.A-expected.A)) >= eps {
			t.Fatalf("#%d a: %e != %e", i, result.A, expected.A)
		}

		if math.Abs(float64(result.B-expected.B)) >= eps {
			t.Fatalf("#%d b: %e != %e", i, result.B, expected.B)
		}
	}
}

func TestLabToRgb(t *testing.T) {
	data := getData()
	for i := 0; i < len(data); i++ {
		input := data[i].lab
		expected := data[i].rgb
		result := OklabToLinearRgb(input)

		if math.Abs(float64(result.R-expected.R)) >= eps {
			t.Fatalf("#%d R: %e != %e", i, result.R, expected.R)
		}

		if math.Abs(float64(result.G-expected.G)) >= eps {
			t.Fatalf("#%d G: %e != %e", i, result.G, expected.G)
		}

		if math.Abs(float64(result.B-expected.B)) >= eps {
			t.Fatalf("#%d B: %e != %e", i, result.B, expected.B)
		}
	}
}

func BenchmarkMath(b *testing.B) {
	data := getData()
	for i := 0; i < b.N; i++ {
		input := data[i%len(data)].rgb
		lab := LinearRgbToOklab(input)
		_ = OklabToLinearRgb(lab)
	}
}
