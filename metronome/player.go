package metronome

import (
	"fmt"
	"os"
	"time"
)

// Output is something that outputs the signals
type Output interface {
	PlayWeak()
	PlayStrong()
}

// Player is a Metronome Player
type Player struct {
	out     Output
	current uint
}

// NewPlayer returns a new Player instance
func NewPlayer(out Output) *Player {
	return &Player{
		out: out,
	}
}

// Reset resets the player to start values
func (p *Player) Reset() {
	p.current = 1
}

// PlayBarUntilSignal plays the given bar until the channel is filled
func (p *Player) PlayBarUntilSignal(bar *Bar, sig chan os.Signal) (err error) {
	return p.PlayBarUntilSignalOrLimit(bar, sig, 0)
}

// PlayBarUntilLimit plays the given bar until the given limit is reached
func (p *Player) PlayBarUntilLimit(bar *Bar, limit uint) (err error) {
	sig := make(chan os.Signal, 0)
	return p.PlayBarUntilSignalOrLimit(bar, sig, limit)
}

// PlayBarUntilSignalOrLimit plays the given bar until either the signal channel is filled the given limit is reached
func (p *Player) PlayBarUntilSignalOrLimit(bar *Bar, sig chan os.Signal, limit uint) (err error) {
	if bar.NoteValue%4 != 0 {
		return fmt.Errorf("Unable to play a bar with a noteValue %q that is not dividable by 4", bar.NoteValue)
	}

	p.Reset()
	d := time.Minute / time.Duration(bar.Tempo*bar.NoteValue/4)
	t := time.NewTicker(d)

	// increase limit by 1 since we start counting at 1 instead of 0
	if limit > 0 {
		limit++
	}

	for {
		select {
		case <-sig:
			t.Stop()
			return

		case <-t.C:
			if p.current == limit {
				return
			}

			if p.current%bar.Beats == 1 {
				go p.out.PlayStrong()

			} else {
				go p.out.PlayWeak()
			}

			p.current++
			break
		default:
		}
	}
}
