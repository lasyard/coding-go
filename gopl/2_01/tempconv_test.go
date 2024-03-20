package tempconv

import (
	"math"
	"testing"
)

func TestCelsiusString(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  string
	}{
		{AbsoluteZeroC, "-273.15ºC"},
		{FreezingC, "0ºC"},
		{BoilingC, "100ºC"},
	}
	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestFahrenheitString(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  string
	}{
		{-459.67, "-459.67ºF"},
		{32, "32ºF"},
		{212, "212ºF"},
	}
	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestKelvinString(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  string
	}{
		{0, "0K"},
		{273.15, "273.15K"},
		{373.15, "373.15K"},
	}
	for _, test := range tests {
		if got := test.input.String(); got != test.want {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestCToF(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Fahrenheit
	}{
		{AbsoluteZeroC, -459.67},
		{FreezingC, 32},
		{BoilingC, 212},
	}
	for _, test := range tests {
		if got := CToF(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{AbsoluteZeroC, 0},
		{FreezingC, 273.15},
		{BoilingC, 373.15},
	}
	for _, test := range tests {
		if got := CToK(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestFToC(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Celsius
	}{
		{-459.67, AbsoluteZeroC},
		{32, FreezingC},
		{212, BoilingC},
	}
	for _, test := range tests {
		if got := FToC(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Kelvin
	}{
		{-459.67, 0},
		{32, 273.15},
		{212, 373.15},
	}
	for _, test := range tests {
		if got := FToK(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{0, AbsoluteZeroC},
		{273.15, FreezingC},
		{373.15, BoilingC},
	}
	for _, test := range tests {
		if got := KToC(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Fahrenheit
	}{
		{0, -459.67},
		{273.15, 32},
		{373.15, 212},
	}
	for _, test := range tests {
		if got := KToF(test.input); math.Abs(float64(got)-float64(test.want)) > 1e-6 {
			t.Errorf("%v = %v, %v wanted", test.input, got, test.want)
		}
	}
}
