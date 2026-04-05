package cmd

import (
	"log"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var projectCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"del"},
	Short:   "Get names of projects",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 0 {
			log.Fatal("this command doesn't take flags")
		}
		run.ShowProjects()
	}}

func init() {
	rootCmd.AddCommand(projectCmd)
}
