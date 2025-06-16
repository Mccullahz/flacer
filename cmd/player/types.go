// this file provides common structs and types used in the player package
package player
import "time"

type AudioMetadata struct {
	Title       string
	Artist      string
	Album       string
	Duration   time.Duration
	SampleRate  int
	BitDepth    int
}

type PlaybackState struct {
	Playing bool
	Position time.Duration
	Duration time.Duration
	Track AudioMetadata
}
