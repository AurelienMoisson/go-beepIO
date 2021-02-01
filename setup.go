package main

import (
	"github.com/gordonklaus/portaudio"
	"math"
)

var harmonics = []float64{0.2, 0.1, 0.03, 0.01}

func Initialize() {
	portaudio.Initialize()
	wave_function = make([]float64, bufferSize)
	for i, _ := range wave_function {
		x := math.Pi * 2 * float64(i) / float64(bufferSize)
		wave_function[i] = 0
		for k, amplitude := range harmonics {
			wave_function[i] += amplitude * math.Sin(x*float64(k+1))
		}
	}
}

func Terminate() {
	portaudio.Terminate()
}
