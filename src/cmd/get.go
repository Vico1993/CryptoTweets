package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Return portfolio special item, otherwise will return everything",
	Long:  `Very small command application to maintain and update my portfolio`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
