// this file contains the logic to manage audio playback (play, pause, rewind, etc.)
// we will need to call decoder.go as a codec to unpack into Pulse-Code Modulation streams that we can send to the DAC. 
package player

import (
)

type Player struct {
}

// initializion
func NewPlayer() *Player {
	
	return &Player{}
}

// load and play an audio file
func (p *Player) Play(path string) error {

	return nil
}

// stop playback
func (p *Player) Stop() error {
	
	return nil
}

// is playing for keeping track of playback state
func (p *Player) IsPlaying() bool {
	
	return false 
}

// output the current file being played for UI
func (p *Player) CurrentFile() string {
	
	return ""
}

