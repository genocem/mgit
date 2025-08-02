package cmd

import (
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var addNamespaceCmd = &cobra.Command{
	Use:   "namespace --name <name>",
	Short: "Add a new namespace",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		run.AddNamespace(name)
	},
}

func init() {
	addNamespaceCmd.Flags().String("name", "", "Repository name")
	addNamespaceCmd.MarkFlagRequired("name")
	addNamespaceCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	orig := addNamespaceCmd.HelpFunc()
	addNamespaceCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		addNamespaceCmd.Flags().Lookup("namespace").Hidden = true
		orig(cmd, args)
	})

	addCmd.AddCommand(addNamespaceCmd)
}
