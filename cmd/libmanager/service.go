package libmanager

import (
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

// internal helper, not exposed
func (s *Service) openFileSelectorInternal() (string, error) {
    if s.ctx == nil {
        return "", errors.New("runtime context is not set")
    }

    filters := []runtime.FileFilter{
        {
            DisplayName: "Lossless Audio",
            Pattern:     "*.flac;*.wav",
        },
    }

    return runtime.OpenFileDialog(s.ctx, runtime.OpenDialogOptions{
        Title:   "Select a music file",
        Filters: filters,
    })
}

// exposed func to frontend
func (s *Service) OpenFileSelector() (string, error) {
    return s.openFileSelectorInternal()
}

func (s *Service) AddMusicFile(filePath string) (Track, error) {
	return s.library.AddTrack(filePath)
}

