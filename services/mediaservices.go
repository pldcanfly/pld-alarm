package services

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
}

func NewMediaState() *MediaState {
	return &MediaState{}
}

func (ms *MediaState) PlayAudio(src string) {
	ms.Type = MediaTypeAudio
	ms.Src = src
	ms.Name = "TEST"
}
