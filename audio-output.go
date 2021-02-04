package main

import (
	"github.com/gordonklaus/portaudio"
	"fmt"
)

type AudioWriter struct {
	bufferChannel chan<- []int32
}

var Wave_function []float64

func audioWriterRoutine(bufferChannel <-chan []int32) {
	buffer := make([]int32, bufferSize)

	stream, err := portaudio.OpenDefaultStream(0, 1, 44100, bufferSize, buffer)
	chk(err)

	chk(stream.Start())

	var receivedBuffer []int32

	for {
		select {
		case receivedBuffer = <-bufferChannel:
			fmt.Println("received buffer: ", buffer)
		default:
		}
		copy(buffer, receivedBuffer)
		stream.Write()
	}
}

func (audioWriter *AudioWriter) WriteBuffer(buffer []int32) {
	audioWriter.bufferChannel <- buffer
}

func (audioWriter *AudioWriter) WriteChord(chord []int) {
	buffer := make([]int32, bufferSize)
	for i, _ := range buffer {
		value := 0.
		for _, k := range chord {
			value += Wave_function[(i*k)%bufferSize]
		}
		buffer[i] = int32(value * maxInt32)
	}
	audioWriter.WriteBuffer(buffer)
}

func StartAudioWriter() AudioWriter {
	commandsChannel := make(chan []int32)
	go audioWriterRoutine(commandsChannel)
	return AudioWriter{commandsChannel}
}
