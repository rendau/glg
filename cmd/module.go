package cmd

import (
	"github.com/rendau/glg/internal/module"
	"github.com/spf13/cobra"
	"log"
)

// moduleCmd represents the module command
var moduleCmd = &cobra.Command{
	Use:   "module",
	Short: "Generates whole module",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		f, err := cmd.Flags().GetString("file")
		if err != nil {
			log.Fatalln("Bad value for file-flag", err)
		}

		n, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatalln("Bad value for name-flag", err)
		}

		module.Run(f, n)
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
	moduleCmd.Flags().StringP("file", "f", "", "Source file name, that contain types")
	moduleCmd.Flags().StringP("name", "n", "", "Module name")

	cobra.MarkFlagRequired(moduleCmd.Flags(), "file")
	cobra.MarkFlagRequired(moduleCmd.Flags(), "name")
}
