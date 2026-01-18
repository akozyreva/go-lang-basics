package main

import "11chapterInterfaces/gadget"

func playList(device gadget.TapeInterface, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	var player gadget.TapeInterface
	player = gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 6"}
	playList(player, mixtape)

	// old  version
	//player := gadget.TapePlayer{}
	//playList(player, mixtape)
	// it won't work! because playList accepts only one type
	//player1 := gadget.TapeRecorder{}
	//playList(player1, mixtape)
	player = gadget.TapeRecorder{}
	playList(player, mixtape)
}
