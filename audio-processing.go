package main

import (
	"github.com/mjibson/go-dsp/fft"
	"math"
	"math/cmplx"
)

const maxInt32 = 2147483647

func GetAmplitudes(buffer []int32) []float64 {
	values := make([]float64, len(buffer))
	for i,v := range(buffer) {
		values[i] = float64(v)/maxInt32
	}
	fftValues := fft.FFTReal(values)
	amplitudes := make([]float64, len(fftValues))
	for i,v := range(fftValues) {
		amplitudes[i] = cmplx.Abs(v)
	}
	return amplitudes
}

func GetProjections(amplitudes []float64, chords[][]int) []float64 {
	magnitudes := make([]float64, len(chords))
	for i, chord := range(chords) {
		magnitudes[i] = 0
		for _, freq := range(chord) {
			magnitudes[i] += amplitudes[freq]
		}
		magnitudes[i] /= math.Sqrt(float64(len(chords)))
	}
	return magnitudes
}
