package media

import "io"

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
