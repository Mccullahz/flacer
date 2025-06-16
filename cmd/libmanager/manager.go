package libmanager

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Library struct {
	Tracks map[string]Track
	Dir    string
}

// create the library and ensure the storage dir exists.
func NewLibrary(path string) (*Library, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}
	return &Library{
		Tracks: make(map[string]Track),
		Dir:    path,
	}, nil
}

// copy an audio file to the library directory and index it
func (l *Library) AddTrack(srcPath string) (Track, error) {
	ext := strings.ToLower(filepath.Ext(srcPath))
	if ext != ".flac" && ext != ".wav" {
		return Track{}, errors.New("unsupported file format")
	}

	id := uuid.NewString()
	dst := filepath.Join(l.Dir, id+ext)

	if err := copyFile(srcPath, dst); err != nil {
		return Track{}, err
	}

	track := Track{
		ID:        id,
		FilePath:  dst,
		Title:     filepath.Base(srcPath),
		Format:    ext[1:], // remove the dot from the extension
		DateAdded: time.Now(),
	}

	l.Tracks[id] = track
	return track, nil
}

// delete a track from the library and disk.
func (l *Library) RemoveTrack(id string) error {
	track, ok := l.Tracks[id]
	if !ok {
		return errors.New("track not found")
	}
	if err := os.Remove(track.FilePath); err != nil {
		return err
	}
	delete(l.Tracks, id)
	return nil
}

// persist the track map to JSON
func (l *Library) Save(jsonPath string) error {
	data, err := json.MarshalIndent(l.Tracks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonPath, data, 0644)
}

// restore the track map from JSON
func (l *Library) Load(jsonPath string) error {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &l.Tracks)
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

