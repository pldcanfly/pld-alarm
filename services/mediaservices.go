package services

import (
	"fmt"

	"github.com/ebitengine/oto/v3"
	"github.com/pldcanfly/pld-alarm/media"
)

const (
	MediaTypeAudio = iota + 1
	MediaTypeYouTube
)

type MediaState struct {
	Duration float64
	Current  float64
	Playing  bool
	Context  *oto.Context
	Player   *oto.Player
	Media    media.Media
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

func (ms *MediaState) GetName() string {
	if !ms.Playing {
		return ""
	}

	return ms.Media.GetName()
}

func (ms *MediaState) PlayAudio(src string) error {
	var err error
	ms.Media, err = media.NewMedia(src)
	if err != nil {
		return fmt.Errorf("Media Error %w", err)

	}
	ms.Playing = true
	ms.Player = ms.Context.NewPlayer(ms.Media.GetStream())
	ms.Player.Play()
	return nil

}

func (ms *MediaState) StopAudio() {
	ms.Playing = false
	ms.Player.Close()
}
