package cmd

import "github.com/spf13/cobra"

func GetRoot() *cobra.Command {
	cmdRoot := &cobra.Command{
		Use: "ggclass",
	}

	cmdRoot.AddCommand(server())
	cmdRoot.AddCommand(migrateUp())
	cmdRoot.AddCommand(migrateDown())
	cmdRoot.AddCommand(migrateRefresh())
	cmdRoot.AddCommand(genError())

	return cmdRoot
}
