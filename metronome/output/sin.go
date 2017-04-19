package output

import "math"

// GenerateSin renders a sin wave with given sampleRate, frequency and number of samples.
func GenerateSin(sampleRate, samples uint, freq float64) []float64 {
	b := make([]float64, samples)
	step := freq / float64(sampleRate)
	var phase float64

	for i := range b {
		b[i] = float64(math.Sin(2 * math.Pi * phase))
		_, phase = math.Modf(phase + step)
	}

	return b
}
