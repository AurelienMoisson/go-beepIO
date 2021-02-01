package main

import (
	"github.com/gordonklaus/portaudio"
	"fmt"
)

var wave_function []float64

func audioWriter(commandsChannel <-chan []int) {
	buffer := make([]int32, bufferSize)

	stream, err := portaudio.OpenDefaultStream(0,1,44100, bufferSize, buffer)
	chk(err)

	chk(stream.Start())

	chord := []int{}
	for {
		select {
		case chord = <-commandsChannel:
			fmt.Println("received chord: ", chord)
		default:
		}
		for i,_ := range(buffer) {
			value := 0.
			for _,k := range(chord) {
				value += wave_function[(i*k) % bufferSize]
			}
			buffer[i] = int32(value)
		}
		stream.Write()
	}
}

func StartAudioWriter() chan<- []int {
	commandsChannel := make(chan []int)
	go audioWriter(commandsChannel)
	return commandsChannel
}
