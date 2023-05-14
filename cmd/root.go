package cmd

import (
	"github.com/spf13/cobra"
	"tour/internal/timer"
	"tour/internal/word"
)

var rootCmd cobra.Command

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(word.WordCmd)
	rootCmd.AddCommand(timer.TimeCmd)
}
