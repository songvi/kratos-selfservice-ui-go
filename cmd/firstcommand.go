package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  "version version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Static version 0.0.1 ")
	},
}
