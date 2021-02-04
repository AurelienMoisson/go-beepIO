package audioIO

import (
	"github.com/gordonklaus/portaudio"
)

type AudioWriter struct {
	stream *portaudio.Stream

	Chord []int
}

func NewAudioWriter() *AudioWriter {
	return &AudioWriter{}	
}

func (writer *AudioWriter) Start() error {
	var err error
	writer.stream, err = portaudio.OpenDefaultStream(0, 1, 44100, bufferSize, writer.processAudio)
	if err != nil {
		return err
	}

	return writer.stream.Start()
}

var Wave_function []float64

func (writer *AudioWriter) processAudio(out []float32) {
	for i, _ := range out {
		out[i] = 0

		for _, k := range writer.Chord {
			out[i] += float32(Wave_function[(i*k)%bufferSize])
		}
	}
}
