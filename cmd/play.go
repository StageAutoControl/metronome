// Copyright © 2017 Alexander Pinnecke <alexander.pinnecke@googlemail.com>

package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/StageAutoControl/metronome/metronome"
	"github.com/StageAutoControl/metronome/metronome/output"
	"github.com/StageAutoControl/metronome/metronome/utils"
	"github.com/spf13/cobra"
)

const (
	outputTypeAudio  = "audio"
	outputTypeStdOut = "stdout"
)

var (
	outputType string
	strongFreq float64
	weakFreq   float64
	limit      uint

	outputTypes = []string{outputTypeStdOut, outputTypeAudio}
)

// playCmd represents the metronome command
var playCmd = &cobra.Command{
	Use:     "play [speed beats noteValue]",
	Example: "play 160 3 4",
	Short:   "Simple metronome with cli and audio output",
	Long:    `A very simple yet flexible metronome using time.Ticker and channels for communication, mainly used for testing some stuff.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Printf("Usage: %s\n", cmd.Use)
			os.Exit(1)
		}

		var speed, beats, noteValue uint64
		var err error

		if speed, err = strconv.ParseUint(args[0], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as speed", args[0]))
		}
		if beats, err = strconv.ParseUint(args[1], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as beats", args[1]))
		}
		if noteValue, err = strconv.ParseUint(args[2], 10, 64); err != nil {
			panic(fmt.Errorf("Unable to parse %q as noteValue", args[2]))
		}

		sig := utils.GetSignal()
		var out metronome.Output

		switch outputType {
		case outputTypeAudio:
			o := output.NewAudioOutput(strongFreq, weakFreq)
			if err := o.Start(); err != nil {
				panic(err)
			}

			defer o.Stop()
			out = o
			break
		case outputTypeStdOut:
			out = output.NewBufferOutput(os.Stdout)
			break
		default:
			panic(fmt.Errorf("Invalid output type %q, valid are %v", outputType, outputTypes))
		}

		m := metronome.NewPlayer(out)
		b := metronome.NewBar(uint(beats), uint(noteValue), uint(speed))

		if err := m.PlayBarUntilSignalOrLimit(b, sig, limit); err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(playCmd)

	playCmd.Flags().StringVarP(&outputType, "output", "o", "audio", fmt.Sprintf("Which output should be used %v", outputTypes))
	playCmd.Flags().Float64Var(&strongFreq, "strongFreq", 1760, "Which frequency should be used to render the sin wave for the strong bar accent click (audio only)")
	playCmd.Flags().Float64Var(&weakFreq, "weakFreq", 1320, "Which frequency should be used to render the sin wave for the weak mediate click (audio only)")
	playCmd.Flags().UintVar(&limit, "limit", 0, "How many clicks to play. Used to limit to a single bar for example")
}
