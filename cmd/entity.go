package cmd

import (
	"log"
	"os"

	"github.com/rendau/glg/internal/cmd/entity"
	"github.com/spf13/cobra"
)

var entityCmd = &cobra.Command{
	Use:   "entity",
	Short: "Generates base entity file",
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

		entity.Run(d, n)

		os.Exit(0)
	},
}

func init() {
	rootCmd.AddCommand(entityCmd)

	entityCmd.Flags().StringP("dir", "d", ".", "Project dir, default is .")
	entityCmd.Flags().StringP("name", "n", "", "Entity name (without 'St' suffix)")

	cobra.MarkFlagRequired(entityCmd.Flags(), "name")
}
