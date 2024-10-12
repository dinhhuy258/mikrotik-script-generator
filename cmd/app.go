package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

func Execute() error {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "application",
		Long:  "application",
		Run:   func(*cobra.Command, []string) {},
	}

	cmd.AddCommand(createServiceCommand())

	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	return err
}
