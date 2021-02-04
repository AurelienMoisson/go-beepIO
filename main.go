package main

import (
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

func sendMessage(commandsChannel chan<- []int, delay int, chord0,chord1 []int, msg []bool) {
	duration := time.Duration(delay) * time.Millisecond
	for _, bit := range msg {
		if bit {
			commandsChannel <- chord0
			time.Sleep(duration)
			commandsChannel <- chord1
			time.Sleep(duration)
		} else {
			commandsChannel <- chord1
			time.Sleep(duration)
			commandsChannel <- chord0
			time.Sleep(duration)
		}
	}
	commandsChannel <- []int{}
}

func getBar(x float64, width int) string {
	barSize := int(x * float64(width))
	if barSize > width {
		barSize = width
	}
	return strings.Repeat("x", barSize) + strings.Repeat(" ", width-barSize)
}

func main() {
	Initialize()
	defer Terminate()

	chordCommandsChannel := StartAudioWriter()
	// go metronome(10, chordCommandsChannel)
	go sendMessage(chordCommandsChannel, 200, []int{8, 10}, []int{9, 11},
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

	audioBufferChannel := StartAudioListener()
	for {
		buffer := <-audioBufferChannel
		line := ""
		amplitudes := GetAmplitudes(buffer)
		for _, v := range amplitudes[10:32] {
			line += getBar(v/2, 8) + "|"
		}
		projections := GetProjections(amplitudes, [][]int{[]int{16, 20, 23}, []int{18, 22, 25}})
		projections_percents := []int{
			int(projections[0] * 100.),
			int(projections[1] * 100.),
		}
		fmt.Println(line, " ", projections_percents)
	}
}
