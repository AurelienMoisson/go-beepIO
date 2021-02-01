package main

import (
	"github.com/gordonklaus/portaudio"
	"math"
)

var harmonics = []float64{1,0.3, 0.12, 0.05, 0.2}

func Initialize() {
	portaudio.Initialize()
	wave_function = make([]float64, bufferSize)
	for i,_ := range(wave_function) {
		x := math.Pi*2*float64(i)/float64(bufferSize)
		wave_function[i] = 0
		for k,amplitude := range(harmonics) {
			wave_function[i] = amplitude*math.Sin(x*float64(k))
		}
	}
}

func Terminate() {
	portaudio.Terminate()
}
