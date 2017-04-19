package utils

import (
	"os"
	"os/signal"
)

// GetSignal returns a channel that get's fired on os kill and interrupt signals
func GetSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, os.Kill)
	return sig
}
