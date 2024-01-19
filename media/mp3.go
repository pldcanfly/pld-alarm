package media

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

type MP3Media struct {
	name   string
	stream io.Reader
	src    string
}

func (m *MP3Media) GetName() string {
	return m.name
}
func (m *MP3Media) GetType() int {
	return MP3
}
func (m *MP3Media) GetStream() io.Reader {
	return m.stream
}
func (m *MP3Media) GetSrc() string {
	return m.src
}

func NewMP3(src string) (*MP3Media, error) {

	file, err := os.ReadFile(fmt.Sprintf("static/%v", src))
	if err != nil {
		return &MP3Media{}, fmt.Errorf("mp3: %w", err)
	}

	fbr := bytes.NewReader(file)
	dec, err := mp3.NewDecoder(fbr)
	if err != nil {
		return &MP3Media{}, fmt.Errorf("mp3: %w", err)
	}

	return &MP3Media{
		name:   "test",
		src:    src,
		stream: dec,
	}, nil
}

func NewRemoteMP3(src string) (*MP3Media, error) {
	file, err := http.Get(src)
	if err != nil {
		return &MP3Media{}, fmt.Errorf("mp3: %w", err)
	}
	defer file.Body.Close()

	dec, err := mp3.NewDecoder(file.Body)
	if err != nil {
		return &MP3Media{}, fmt.Errorf("mp3: %w", err)
	}

	return &MP3Media{
		name:   "test",
		src:    src,
		stream: dec,
	}, nil
}
