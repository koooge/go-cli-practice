package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var name string

func NewCmdHey() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hey",
		Short: "hey command",
		Long:  `hey command`,
		Run: func(cmd *cobra.Command, args []string) {
			if name != "" {
				cmd.Println("Hey " + name + "!")
			} else {
				if len(args) == 0 {
					cmd.Println("Hey " + os.Getenv("USER") + "!")
				} else {
					cmd.Println("Hey " + args[0] + "!")
				}
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "./go-cli-practice hey -n <name>")

	return cmd
}
