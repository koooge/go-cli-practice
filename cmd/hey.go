package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var heyCmd = &cobra.Command{
	Use:   "hey",
	Short: "hey command",
	Long:  `hey command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hey " + os.Getenv("USER") + "!")
	},
}
