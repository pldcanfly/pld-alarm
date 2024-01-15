package services

import (
	"bytes"
	"fmt"
	"os"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
)

const (
	MediaTypeAudio = iota + 1
	MediaTypeYouTube
)

type MediaState struct {
	Type     int
	Duration float64
	Current  float64
	Playing  bool
	Src      string
	Name     string
	Context  *oto.Context
	Player   *oto.Player
}

func NewMediaState() *MediaState {
	ctx, ch, err := oto.NewContext(&oto.NewContextOptions{
		SampleRate:   44100,
		ChannelCount: 2,
		Format:       oto.FormatSignedInt16LE,
	})
	<-ch
	if err != nil {
		panic(fmt.Sprintf("oto ctx failed: %v", err))
	}
	return &MediaState{
		Context: ctx,
	}
}

func (ms *MediaState) PlayAudio(src string) {
	ms.Type = MediaTypeAudio
	ms.Src = src
	ms.Name = "TEST"
	ms.Playing = true

	file, err := os.ReadFile(fmt.Sprintf("static/%v", src))
	if err != nil {
		fmt.Println("could not read mp3")
		return
	}

	fbr := bytes.NewReader(file)
	dec, err := mp3.NewDecoder(fbr)
	if err != nil {
		fmt.Println("decoder failed")
		return
	}

	ms.Player = ms.Context.NewPlayer(dec)
	ms.Player.Play()

}

func (ms *MediaState) StopAudio() {
	ms.Type = 0
	ms.Src = ""
	ms.Name = ""
	ms.Playing = false

	ms.Player.Close()
}
