package services

type MediaState struct {
	Type     int
	Duration int
	Current  int
	Playing  bool
	Src      string
}

func NewMediaState() *MediaState {
	return &MediaState{}
}
