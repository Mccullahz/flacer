// the job of the decoder is take in flac or wav files and deode them into PCM streams that should be easily handled by player.go
package player

import (
	"io"
	"os"
	"fmt"
)

// define the methods required for decoding audio files
type Decoder interface {
	Decode(r io.Reader) ([]byte, error) // reads from the provided reader and returns PCM data
	Close() error                       // releases any resources held by the decoder
}

// create a new Decoder based on the file extension
func NewDecoder(filePath string) (Decoder, error) {
	ext := getFileExtension(filePath)
	switch ext {
	case ".wav":
		return NewWavDecoder(filePath)
	case ".flac":
		return NewFlacDecoder(filePath)
	default:
		return nil, fmt.Errorf("unsupported file format: %s", ext)
	}
}

// return the file extension of the given file path
func getFileExtension(filePath string) string {
	if len(filePath) < 4 {
		return ""
	}
	return filePath[len(filePath)-4:]
}

// decode the audio file at the given path and return PCM data
func DecodeFile(filePath string) ([]byte, error) {
	decoder, err := NewDecoder(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to create decoder: %w", err)
	}
	defer decoder.Close()

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	pcmData, err := decoder.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("failed to decode file: %w", err)
	}

	return pcmData, nil
}

// release any resources held by the decoder
func (d *Decoder) Close() error {
	if d == nil {
		return nil
	}
	if closer, ok := d.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

