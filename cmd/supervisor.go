package cmd

import (
	"log/slog"

	"github.com/spf13/cobra"
)

var supervisorCmd = &cobra.Command{
	Use:     "supervisor",
	Aliases: []string{"super", "su"},
	Short:   "Monitor, control and configure the ApexOS Supervisor",
	Long: `
The ApexOS Supervisor is the heart of the ApexOS system.
It manages your ApexOS Core, Operating System, and all the apps.
It even manages itself! This series of command give you control over the
ApexOS Supervisor.`,
	Example: `
  apex supervisor reload
  apex supervisor update
  apex supervisor logs`,
}

func init() {
	slog.Debug("Init supervisor")
	rootCmd.AddCommand(supervisorCmd)
}
