package cmd

import (
	"github.com/rendau/glg/internal/cmd/module"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// moduleCmd represents the module command
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	moduleCmd.Flags().StringP("dir", "d", ".", "Project dir, default is .")
	moduleCmd.Flags().StringP("name", "n", "", "Entity name (without 'St' suffix)")

	cobra.MarkFlagRequired(moduleCmd.Flags(), "name")
}
