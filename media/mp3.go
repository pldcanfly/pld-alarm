package media

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

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
	var dec io.Reader
	var err error
	if strings.HasPrefix(src, "http://") || strings.HasPrefix(src, "https://") {
		dec, err = remoteMP3(src)
	} else {
		dec, err = localMP3(src)
	}

	if err != nil {
		return &MP3Media{}, err
	}

	return &MP3Media{
		name:   "test",
		src:    src,
		stream: dec,
	}, nil
}

func localMP3(src string) (io.Reader, error) {
	file, err := os.ReadFile(fmt.Sprintf("static/%v", src))
	if err != nil {
		return nil, fmt.Errorf("mp3: %w", err)
	}

	fbr := bytes.NewReader(file)
	dec, err := mp3.NewDecoder(fbr)
	if err != nil {
		return nil, fmt.Errorf("mp3: %w", err)
	}

	return dec, nil
}

func remoteMP3(src string) (io.Reader, error) {
	file, err := http.Get(src)
	if err != nil {
		return nil, fmt.Errorf("remote mp3: %w", err)
	}
	defer file.Body.Close()

	dec, err := mp3.NewDecoder(file.Body)
	if err != nil {
		return nil, fmt.Errorf("remote mp3: %w", err)
	}
	return dec, nil

}
