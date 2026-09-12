package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var audioDefaultCmd = &cobra.Command{
	Use:     "default",
	Aliases: []string{"def", "de", "standard", "std"},
	Short:   "Set default input/output audio device.",
	Long: `
Set the default input/output audio device of your ApexOS system.
`,
	Example: `
	apex audio default input --name "..."
	apex audio default output --name "..."
`,
}

func init() {
	slog.Debug("Init audio default")

	audioCmd.AddCommand(audioDefaultCmd)
}
