#include <cstdio>
#include <cmath>
#include <cassert>
#include "oklab.cpp"

double srgbToLinear(double cbNumber) {
    assert(cbNumber >= 0.0);
    assert(cbNumber <= 255.0);

    double c = cbNumber / 255.0;

    assert(c >= 0.0);
    assert(c <= 1.0);

    double a = 0.055;
    if (c <= 0.04045) {
        return c / 12.92;
    } else {
        double result = pow((c + a) / (1 + a), 2.4);
        assert(result >= 0.0);
        assert(result <= 1.0);
        return result;
    }
}

void process(int r, int g, int b) {
    double lR = srgbToLinear(r);
    double lG = srgbToLinear(g);
    double lB = srgbToLinear(b);

    printf("R: %e, G: %e, B: %e\n", lR, lG, lB);

    RGB input;
    input.r = lR;
    input.g = lG;
    input.b = lB;
    Lab output = linear_srgb_to_oklab(input);

    RGB convertedBack = oklab_to_linear_srgb(output);
    assert(abs(convertedBack.r - input.r) < 0.000001);
    assert(abs(convertedBack.g - input.g) < 0.000001);
    assert(abs(convertedBack.b - input.b) < 0.000001);

    printf("L: %e, a %e, b %e\n\n", output.L, output.a, output.b);
}

int main() {
    printf("\n");

    process(255, 0, 0);
    process(0, 255, 0);
    process(0, 0, 255);

    process(255, 255, 255);
    process(0, 0, 0);

    process(136, 212, 130);
    process(255, 159, 56);
    process(196, 190, 166);

    process(255, 188, 255);
    process(255, 18, 125);
    process(255, 0, 75);

    return 0;
}
