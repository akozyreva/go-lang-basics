package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopping")
}

type TapeRecorder struct {
	Microphones int
}

// Play there're duplicates of TapePlayer
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopping")
}

type TapeInterface interface {
	Play(song string)
	Stop()
}
