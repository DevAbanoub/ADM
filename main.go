package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "sidm",
		Short:        "Smart Internet Download Manager!",
		SilenceUsage: true,
	}
	cmd.AddCommand(downloadUrlCmd())
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func downloadUrlCmd() *cobra.Command {
	return &cobra.Command{
		Use: "get",
		RunE: func(cmd *cobra.Command, args []string) error {
			// now := time.Now()
			// prettyTime := now.Format(time.RubyDate)
			cmd.Println("Downloading...")
			return nil
		},
	}
}
