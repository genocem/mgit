package cmd

import (
	"log"
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var namespaceCmd = &cobra.Command{
	Use:     "namespace",
	Aliases: []string{"del"},
	Short:   "Get names of namespaces",
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flags().NFlag() != 0 {
			log.Fatal("this command doesn't take flags")
		}
		run.ShowNamespaces()
	}}

func init() {
	rootCmd.AddCommand(namespaceCmd)
}
