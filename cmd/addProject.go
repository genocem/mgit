package cmd

import (
	"mgit/internal/run"

	"github.com/spf13/cobra"
)

var addProjectCmd = &cobra.Command{
	Use:   "project --name <name>",
	Short: "Add a new project",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		run.AddProject(name)
	},
}

func init() {
	addProjectCmd.Flags().String("name", "", "Repository name")
	addProjectCmd.MarkFlagRequired("name")
	addProjectCmd.RegisterFlagCompletionFunc("name", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveNoSpace | cobra.ShellCompDirectiveNoFileComp
	})

	orig := addProjectCmd.HelpFunc()
	addProjectCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		addProjectCmd.Flags().Lookup("project").Hidden = true
		orig(cmd, args)
	})

	addCmd.AddCommand(addProjectCmd)
}
