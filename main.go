package main

import (
	"github.com/AurelienMoisson/go-beepIO/audioIO"
	"fmt"
	"time"
	"strings"
)

func metronome(delay int, commandsChannel chan<- []int) {
	fmt.Println("starting metronome")
	for {
		for i := 0; i < delay; i++ {
			commandsChannel <- []int{16, 20, 23}
		}
		for i := 0; i < delay; i++ {
			commandsChannel <- []int{18, 22, 25}
		}
	}
}

func sendMessage(audioWriter *audioIO.AudioWriter, delay int, chord0,chord1 []int, msg []bool) {
	duration := time.Duration(delay) * time.Millisecond
	for _, bit := range msg {
		if bit {
			audioWriter.Chord = chord0
		} else {
			audioWriter.Chord = chord1
		}
		time.Sleep(duration)
	}
	audioWriter.Chord = []int{}
}

func getBar(x float64, width int) string {
	barSize := int(x * float64(width))
	if barSize > width {
		barSize = width
	}
	return strings.Repeat("x", barSize) + strings.Repeat(" ", width-barSize)
}

func main() {
	audioIO.Initialize()
	defer audioIO.Terminate()

	audioWriter := audioIO.NewAudioWriter()
	if err := audioWriter.Start(); err != nil {
		panic(err)
	}

	// go metronome(10, chordCommandsChannel)
	go sendMessage(audioWriter, 200, []int{8, 10}, []int{9, 11},
		[]bool{
			true,
			true,
			false,
			true,
			false,
			false,
			false,
			true,
			true,
			false,
			false,
			false,
			false,
			true,
			false,
			true,
			true,
		},
	)

	audioListener := audioIO.NewAudioListener(processAudio)
	if err := audioListener.Start(); err != nil {
		panic(err)
	}

	<-make(chan struct{})
}

func processAudio(buffer []float64) {
	line := ""

	amplitudes := audioIO.GetAmplitudes(buffer)
	for _, v := range amplitudes[8:22] {
		line += getBar(v/2, 8) + "|"
	}

	projections := audioIO.GetProjections(amplitudes, [][]int{[]int{8,10}, []int{9,11}})
	projections_percents := []int{
		int(projections[0] * 100.),
		int(projections[1] * 100.),
	}

	fmt.Println(line, " ", projections_percents)
}
