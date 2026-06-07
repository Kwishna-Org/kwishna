package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "kwishna",
	Short: "a lightweight reverse-proxy protection layer that filters automated traffic.",
}

func Execute() {
	if err := initCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// TODO
}