package main

import (
	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate    = 44100
	bufferSize    = 1024
	baseFrequency = sampleRate / bufferSize
)

type AudioListener struct {
	bufferChannel <-chan []int32
}

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

func (audioListener *AudioListener) GetAudioBuffer() []int32 {
	return <-audioListener.bufferChannel
}

func StartAudioListener() AudioListener {
	bufferChannel := make(chan []int32)
	go audioListener(bufferChannel)
	return AudioListener{bufferChannel}
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
