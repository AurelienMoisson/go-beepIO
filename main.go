package main

import (
	"fmt"
	"strings"
)

func metronome(delay int, commandsChannel chan<- []int) {
	fmt.Println("starting metronome")
	for {
		for i := 0; i < delay; i++ {
			commandsChannel <- []int{16, 20, 23}
		}
		for i := 0; i < delay; i++ {
			commandsChannel <- []int{}
		}
	}
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
	go metronome(10, chordCommandsChannel)

	audioBufferChannel := StartAudioListener()
	for {
		buffer := <-audioBufferChannel
		line := ""
		amplitudes := GetAmplitudes(buffer)
		for _, v := range amplitudes[16:32] {
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
