package cmd

import (
	"log"
	"os"

	"github.com/rendau/glg/internal/cmd/module"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Generates whole module",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		d, err := cmd.Flags().GetString("dir")
		if err != nil {
			log.Fatalln("Bad value for dir-flag", err)
		}

		n, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalln("Bad value for name-flag", err)
		}

		module.Run(d, n)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(moduleCmd)

	moduleCmd.Flags().StringP("dir", "d", ".", "Project dir, default is .")
	moduleCmd.Flags().StringP("name", "n", "", "Entity name (without 'St' suffix)")

	cobra.MarkFlagRequired(moduleCmd.Flags(), "name")
}
