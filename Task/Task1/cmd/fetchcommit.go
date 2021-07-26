package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func Execute(command *cobra.Command) error {
	if err := command.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return err
	}
	return nil
}
