package media

import (
	"io"
	"strings"
)

type Media interface {
	GetName() string
	GetStream() io.Reader
	GetType() int
	GetSrc() string
}

const (
	MP3 = iota + 1
	FLAC
)

func NewMedia(src string) (Media, error) {
	if strings.HasSuffix(src, ".mp3") {
		return NewMP3(src)
	}

	return nil, nil
}
