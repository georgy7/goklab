package goklab

import (
	"math"
	"testing"
)

const eps = 1e-6

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
			Lab{8.664396e-01, -2.338876e-01, 1.794985e-01},
		},
		{
			RGB{0.000000e+00, 0.000000e+00, 1.000000e+00},
			Lab{4.520137e-01, -3.245698e-02, -3.115281e-01},
		},
		{
			RGB{1.000000e+00, 1.000000e+00, 1.000000e+00},
			Lab{1.000000e+00, 8.095286e-11, 3.727391e-08},
		},
		{
			RGB{0.000000e+00, 0.000000e+00, 0.000000e+00},
			Lab{0.000000e+00, 0.000000e+00, 0.000000e+00},
		},
		{
			RGB{2.462013e-01, 6.583748e-01, 2.232280e-01},
			Lab{8.002463e-01, -1.078114e-01, 8.255051e-02},
		},
		{
			RGB{1.000000e+00, 3.467041e-01, 3.954624e-02},
			Lab{7.842250e-01, 7.268992e-02, 1.413031e-01},
		},
		{
			RGB{5.520114e-01, 5.149177e-01, 3.813260e-01},
			Lab{7.999292e-01, -3.195469e-03, 3.335571e-02},
		},
		{
			RGB{1.000000e+00, 5.028865e-01, 1.000000e+00},
			Lab{8.774487e-01, 9.605264e-02, -6.356862e-02},
		},
		{
			RGB{1.000000e+00, 6.048833e-03, 2.050787e-01},
			Lab{6.477489e-01, 2.548900e-01, 1.702712e-02},
		},
		{
			RGB{1.000000e+00, 0.000000e+00, 7.036010e-02},
			Lab{6.337077e-01, 2.412873e-01, 7.909187e-02},
		},
	}
}

func TestItIsReversible(t *testing.T) {
	data := getData()
	for i := 0; i < len(data); i++ {
		input := data[i].rgb
		lab := LinearRgbToOklab(input)
		result := OklabToLinearRgb(lab)

		if math.Abs(result.R-input.R) >= eps {
			t.Fatalf("#%d R: %e != %e", i, result.R, input.R)
		}

		if math.Abs(result.G-input.G) >= eps {
			t.Fatalf("#%d G: %e != %e", i, result.G, input.G)
		}

		if math.Abs(result.B-input.B) >= eps {
			t.Fatalf("#%d B: %e != %e", i, result.B, input.B)
		}
	}
}

func TestRgbToLab(t *testing.T) {
	data := getData()
	for i := 0; i < len(data); i++ {
		input := data[i].rgb
		expected := data[i].lab
		result := LinearRgbToOklab(input)

		if math.Abs(result.L-expected.L) >= eps {
			t.Fatalf("#%d L: %e != %e", i, result.L, expected.L)
		}

		if math.Abs(result.A-expected.A) >= eps {
			t.Fatalf("#%d a: %e != %e", i, result.A, expected.A)
		}

		if math.Abs(result.B-expected.B) >= eps {
			t.Fatalf("#%d b: %e != %e", i, result.B, expected.B)
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
