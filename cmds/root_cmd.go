package cmds

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var RootCmd = &cobra.Command{
	Use:   "devops",
	Short: "hello world~",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello world~")
	},
}

func init() {
	RootCmd.AddCommand(runCMD, execCMD)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
