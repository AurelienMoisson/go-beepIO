package main

import (
	"time"
	"fmt"
	"strings"
)

func metronome(delay int, commandsChannel chan<- []int) {
	for {
		for i:=0; i<delay; i++ {
			commandsChannel <- []int{16, 20, 23}
		}
		for i:=0; i<delay; i++ {
			commandsChannel <- []int{}
		}
	}
}

func getBar(x float64, width int) string {
	barSize := int(x*float64(width))
	return strings.Repeat("x", barSize) + strings.Repeat(" ", width-barSize)
}


func main() {
	Initialize()
	defer Terminate()

	/*
	chordCommandsChannel := StartAudioWriter()
	go metronome(10, chordCommandsChannel)
	*/

	time.Sleep(10)
	audioBufferChannel := StartAudioListener()
	for {
		buffer := <-audioBufferChannel
		//fmt.Println(buffer)
		line := ""
		for _, v := range(GetAmplitudes(buffer)[16:32]) {
			line += getBar(v, 8) + "|"
		}
		fmt.Println(line)
	}
}