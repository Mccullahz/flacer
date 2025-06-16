package libmanager

import "time"

type Track struct {
	ID         string    `json:"id"`
	FilePath   string    `json:"filePath"`
	Title      string    `json:"title"`
	Format     string    `json:"format"`
	Album      string    `json:"album"`
	Original   string    `json:"original"`
	DateAdded  time.Time `json:"dateAdded"`
}
