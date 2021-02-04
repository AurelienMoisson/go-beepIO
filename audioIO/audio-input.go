package audioIO

import (
	"github.com/gordonklaus/portaudio"
)

const (
	sampleRate    = 44100
	bufferSize    = 1024
	baseFrequency = sampleRate / bufferSize
)

type AudioListener struct {
	stream *portaudio.Stream
	callback func([]float64)
}

func NewAudioListener(callback func([]float64)) *AudioListener {
	return &AudioListener{
		callback: callback,
	}
}

func (listener *AudioListener) Start() error {
	var err error
	listener.stream, err = portaudio.OpenDefaultStream(1, 0, sampleRate, bufferSize, listener.processAudio)
	if err != nil {
		return err
	}

	return listener.stream.Start()
}


func (listener *AudioListener) Stop() error {
	return listener.stream.Close()
}

func (listener *AudioListener) processAudio(in []float32){
	buffer := make([]float64, len(in))

	for i, sample := range in {
		buffer[i] = float64(sample)
	}

	listener.callback(buffer)
}
