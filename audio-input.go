package main

import (
	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate    = 44100
	bufferSize    = 1024
	baseFrequency = sampleRate / bufferSize
)

func audioListener(bufferChannel chan<- []int32) {

	in := make([]int32, bufferSize)
	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	chk(err)
	defer stream.Close()

	chk(stream.Start())
	for {
		chk(stream.Read())
		select {
		case bufferChannel <- in:
		default:
		}
	}
}

func StartAudioListener() <-chan []int32 {
	bufferChannel := make(chan []int32)
	go audioListener(bufferChannel)
	return bufferChannel
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
