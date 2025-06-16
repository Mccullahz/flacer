package libmanager

import (
	"os"
	"path/filepath"
	"strings"
	"errors"
	"context"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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
	return &Service{library: lib}
}

// frontend calls this to select a folder
func (s *Service) OpenDirectorySelector() (string, error) {
	return runtime.OpenDirectoryDialog(s.ctx, runtime.OpenDialogOptions{
		Title: "Select a music folder",
	})
}

// frontend calls this to add tracks from folder
func (s *Service) AddMusicFolder(folderPath string) ([]Track, error) {
	if folderPath == "" {
		return nil, errors.New("no folder selected")
	}

	albumName := filepath.Base(folderPath) // use folder name as album name
	var tracks []Track

	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(path))
		if ext == ".flac" || ext == ".wav" {
			track, err := s.library.AddTrackToAlbum(path, albumName)
			if err == nil {
				tracks = append(tracks, track)
			}
		}
		return nil
	})
	// scan for album art
	files, err := os.ReadDir(folderPath)
	if err == nil {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".jpg" || ext == ".jpeg" || ext == ".png" {
				srcCover := filepath.Join(folderPath, file.Name())
				destCover := filepath.Join(s.library.Dir, albumName, "cover"+ext)
				_ = copyFile(srcCover, destCover)
				break // only copy the first valid cover found
			}
		}
	}


	if err != nil {
		return nil, err
	}
	return tracks, nil
}

