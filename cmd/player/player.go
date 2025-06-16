// this file contains the logic to manage audio playback (play, pause, rewind, etc.)
// we will need to call decoder.go as a codec to unpack into Pulse-Code Modulation streams that we can send to the DAC. 
package player

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-audio/wav"
	"github.com/hajimehoshi/oto/v2"

)

type Player struct {
	context     *oto.Context
	player      oto.Player
	isPlaying   bool
	currentFile string
	mutex       sync.Mutex
}

// initializion
func NewPlayer() *Player {
	return &Player{}
}

// load and play a WAV file (currently using go-audio, but will be building the custom decoder for this)
func (p *Player) Play(path string) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.isPlaying {
		p.Stop()
	}

	f, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	decoder := wav.NewDecoder(f)

	if !decoder.IsValidFile() {
		return fmt.Errorf("invalid WAV file")
	}

	buf, err := decoder.FullPCMBuffer()
	if err != nil {
		return fmt.Errorf("failed to decode WAV: %w", err)
	}

	// oto.Context
	if p.context == nil {
		ctx, ready, err := oto.NewContext(decoder.SampleRate, decoder.NumChans, 2) // 2 = 16-bit
		if err != nil {
			return fmt.Errorf("oto context error: %w", err)
		}
		<-ready
		p.context = ctx
	}

	p.player = p.context.NewPlayer()
	p.isPlaying = true
	p.currentFile = path

	go func() {
		defer p.player.Close()
		_, _ = p.player.Write(buf.AsBytes())
		p.mutex.Lock()
		p.isPlaying = false
		p.mutex.Unlock()
	}()

	return nil
}

// stop playback
func (p *Player) Stop() error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.player != nil {
		p.player.Close()
		p.player = nil
	}
	p.isPlaying = false
	return nil
}


func (p *Player) IsPlaying() bool {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.isPlaying
}

func (p *Player) CurrentFile() string {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.currentFile
}

