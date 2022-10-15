package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Short: "Simple archiver",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Execute err:%s", err)
	}
}
