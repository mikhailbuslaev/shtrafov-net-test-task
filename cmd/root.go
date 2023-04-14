package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Cmd ...
var Cmd = &cobra.Command{
	Use:          "sntt",
	Short:        "sntt service",
	SilenceUsage: true,
}

// Execute ...
func Execute() {
	err := Cmd.Execute()

	if err != nil {
		log.Fatal().Err(err).Msg("got err while running")
	}
}

func init() {
	Cmd.AddCommand(ServerCmd)
}
