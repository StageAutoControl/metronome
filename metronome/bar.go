package metronome

import "fmt"

// Bar represents a bar in music with the number of beats, the beat value and the speed
type Bar struct {
	Beats     uint
	NoteValue uint
	Tempo     uint
}

// NewBar returns a new bar with given values
func NewBar(beats, noteValue, tempo uint) *Bar {
	return &Bar{
		Beats:     beats,
		NoteValue: noteValue,
		Tempo:     tempo,
	}
}

// String returns the string representation of a bar
func (b *Bar) String() string {
	return fmt.Sprintf("%d/%d @ %d BPM", b.Beats, b.NoteValue, b.Tempo)
}
