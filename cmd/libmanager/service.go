package libmanager

import (
	"os"
	"fmt"
	"time"
	"path/filepath"
	"strings"
	"errors"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// track struct from types.go
type Track struct {
	ID        string    `json:"id"`
	FilePath  string    `json:"filePath"`
	Title     string    `json:"title"`
	Format    string    `json:"format"`
	Album     string    `json:"album"`
	Artist    string    `json:"artist"`
	Original  string    `json:"original"`
	DateAdded time.Time `json:"dateAdded"`
	CoverPath string    `json:"coverPath"`
}


type Service struct {
	library *Library
	ctx     context.Context
}

// called from main app to set the runtime context
func (s *Service) SetContext(ctx context.Context) {
    s.ctx = ctx
}

func NewService() *Service {
	lib, _ := NewLibrary("./library")
	lib.ScanLibrary() // load existing tracks
	return &Service{library: lib}
}

// frontend calls this to select a folder
func (s *Service) OpenDirectorySelector() (string, error) {
	return runtime.OpenDirectoryDialog(s.ctx, runtime.OpenDialogOptions{
		Title: "Select a music folder",
	})
}

// frontend calls this to add tracks from folder
// sorry to future me, this function is a clusterfuck
func (s *Service) AddMusicFolder(folderPath string) ([]Track, error) {
	if folderPath == "" {
		return nil, errors.New("no folder selected")
	}

	folderName := filepath.Base(folderPath)
	parsedAlbumName := parseAlbumName(folderName)
	albumDir := filepath.Join(s.library.Dir, parsedAlbumName)

	if err := os.MkdirAll(albumDir, 0755); err != nil {
		return nil, err
	}

	var tracks []Track
	var copiedCover bool
	var finalCoverPath string

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))

		if ext == ".flac" || ext == ".wav" {
			track, err := s.library.AddTrackToAlbum(path, folderName)
			if err == nil {
				tracks = append(tracks, track)
			}
		} else if !copiedCover && (ext == ".jpg" || ext == ".jpeg" || ext == ".png") {
			dest := filepath.Join(albumDir, "cover"+ext)
			if err := copyFile(path, dest); err == nil {
				finalCoverPath = dest
				copiedCover = true
			}
		}
		return nil
	})
	s.library.ScanLibrary() // refresh library after adding tracks
	if err != nil {
		return nil, err
	}

	// set cover path for each track and update the in-memory library
	if finalCoverPath != "" {
		absCoverPath, _ := filepath.Abs(finalCoverPath)
		for i, track := range tracks {
			track.CoverPath = absCoverPath
			s.library.Tracks[track.ID] = track
			tracks[i] = track // update slice
		}
		// Save updated library to disk
		_ = s.library.Save(filepath.Join(s.library.Dir, "library.json"))
	}

	return tracks, nil
}

// return all tracks in the library
func (s *Service) GetAllTracks() ([]Track, error) {
	tracks := make([]Track, 0, len(s.library.Tracks))
	for _, track := range s.library.Tracks {
		tracks = append(tracks, track)
	}
	return tracks, nil
}

func (s *Service) RescanLibrary() ([]Track, error) {
	if err := s.library.ScanLibrary(); err != nil {
		return nil, err
	}
	fmt.Println("Rescanned library, total tracks:", len(s.library.Tracks))
	return s.GetAllTracks()
}

