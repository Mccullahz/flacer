package libmanager

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
    	"github.com/dhowden/tag"
	"github.com/google/uuid"
)


type Library struct {
	Tracks map[string]Track
	Dir    string
}

// create the library and ensure the storage directory exists
func NewLibrary(path string) (*Library, error) {
	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	lib := &Library{
		Tracks: make(map[string]Track),
		Dir:    path,
	}

	// try to load previous data
	_ = lib.Load(filepath.Join(path, "library.json"))

	return lib, nil
}

// copy a file to an album folder
func (l *Library) AddTrackToAlbum(srcPath string, albumName string) (Track, error) {
	ext := strings.ToLower(filepath.Ext(srcPath))
	if ext != ".flac" && ext != ".wav" {
		return Track{}, errors.New("unsupported file format")
	}

	originalName := filepath.Base(srcPath)
	albumDir := filepath.Join(l.Dir, albumName)
	if err := os.MkdirAll(albumDir, 0755); err != nil {
		return Track{}, err
	}

	dst := filepath.Join(albumDir, originalName)
	if err := copyFile(srcPath, dst); err != nil {
		return Track{}, err
	}

	id := uuid.NewString()

	var artist string = "Unknown Artist"
    	f, err := os.Open(dst)
    	if err == nil {
        	defer f.Close()
        	metadata, err := tag.ReadFrom(f)
        	if err == nil {
        	    if metadata.Artist() != "" {
	                artist = metadata.Artist()
	            }
        	}
    	}

	track := Track{
		ID:        id,
		FilePath:  dst,
		Original:  originalName,
		Title:     strings.TrimSuffix(originalName, ext),
		Format:    ext[1:], // remove dot
		Album:     albumName,
		Artist:    artist, // lsp throwing error, but this works lol
		DateAdded: time.Now(),
	}

	l.Tracks[id] = track
	// save the updated lib to json for persistence
	if err := l.Save(filepath.Join(l.Dir, "library.json")); err != nil {
		return track, err
	}

	return track, nil
}

// delete a track from the library
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

// persist the track map to a JSON file
func (l *Library) Save(jsonPath string) error {
	data, err := json.MarshalIndent(l.Tracks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonPath, data, 0644)
}

// restore the track map from a JSON file
func (l *Library) Load(jsonPath string) error {
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &l.Tracks)
}

// internal copy utility
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

// used to rescan the lib. recursively walks the library directory can be ran on startup
func (l *Library) ScanLibrary() error {
	err := filepath.Walk(l.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		if ext != ".flac" && ext != ".wav" {
			return nil
		}

		// !duplicates
		for _, track := range l.Tracks {
			if track.FilePath == path {
				return nil
			}
		}
		albumName := filepath.Base(filepath.Dir(path))
		// autosaves
		_, err = l.AddTrackToAlbum(path, albumName)
		return err
	})
	return err
}

