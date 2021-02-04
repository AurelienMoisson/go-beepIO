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
	buffer []int32
}

func NewAudioListener() *AudioListener {
	return &AudioListener{
		buffer: make([]int32, bufferSize),
	}
}

func (listener *AudioListener) Start() error {
	var err error
	listener.stream, err = portaudio.OpenDefaultStream(1, 0, 44100, len(listener.buffer), listener.buffer)
	if err != nil {
		return err
	}

	return listener.stream.Start()
}


func (listener *AudioListener) Stop() error {
	return listener.stream.Close()
}

func (listener *AudioListener) Next() ([]int32, error) {
	err := listener.stream.Read()
	if err != nil {
		return nil, err
	}

	return listener.buffer, nil
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
