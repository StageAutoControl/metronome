package output

import (
	"errors"

	"github.com/gordonklaus/portaudio"
)

const sampleRate uint = 44100
const numSamples uint = 1000

// AudioOutput is a output stream to audio
type AudioOutput struct {
	*portaudio.Stream
	strong, weak           chan struct{}
	strongSound, weakSound []float64
}

// NewAudioOutput returns a new AudioOutput instance with default values
func NewAudioOutput(strongFreq, weakFreq float64) *AudioOutput {
	return &AudioOutput{
		Stream:      nil,
		strong:      make(chan struct{}, 1),
		weak:        make(chan struct{}, 1),
		strongSound: GenerateSin(sampleRate, numSamples, strongFreq),
		weakSound:   GenerateSin(sampleRate, numSamples, weakFreq),
	}
}

// Start starts the output channel
func (o *AudioOutput) Start() (err error) {
	if err = portaudio.Initialize(); err != nil {
		return
	}

	o.Stream, err = portaudio.OpenDefaultStream(0, 1, float64(sampleRate), 0, o.processAudio)
	if err != nil {
		return
	}

	return o.Stream.Start()
}

// Stop stops the audio output
func (o *AudioOutput) Stop() error {
	// make sure to terminate the audio device and delete the stream!
	defer portaudio.Terminate()
	defer func() {
		o.Stream = nil
	}()

	err := o.Stream.Stop()
	if err != nil {
		return err
	}

	return o.Stream.Close()
}

func (o *AudioOutput) processAudio(b []float32) {
	data := make([]float64, len(b))

	select {
	case <-o.strong:
		data = o.strongSound[:len(b)]
	case <-o.weak:
		data = o.weakSound[:len(b)]
	default:
	}

	for i := range b {
		b[i] = float32(data[i] * 2)
	}
}

// PlayStrong plays a accent note for full bars
func (o *AudioOutput) PlayStrong() {
	if o.Stream == nil {
		panic(errors.New("AudioOutput is not started yet or terminated"))
	}

	o.strong <- struct{}{}
}

// PlayWeak plays a mediate sound sample for 4ths etc.
func (o *AudioOutput) PlayWeak() {
	if o.Stream == nil {
		panic(errors.New("AudioOutput is not started yet or terminated"))
	}

	o.weak <- struct{}{}
}
