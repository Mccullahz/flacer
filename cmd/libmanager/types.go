package libmanager

import "time"

type Track struct {
	ID         string        `json:"id"`
	FilePath   string        `json:"file_path"`
	Title      string        `json:"title"`
	Format     string        `json:"format"` // flac or wav
	Duration   time.Duration `json:"duration,omitempty"` // optional, but i would like this to be filled in by the decoder
	DateAdded  time.Time     `json:"date_added"`
}

