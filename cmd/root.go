package cmd

import (
	"errors"
	"os"
	"regexp"

	"github.com/bersen66/grep/pkg/filter"
	"github.com/spf13/cobra"
)

var fConfig = &filter.Config{}

var rootCmd = &cobra.Command{
	Use:   "grep [flags] [pattern]",
	Short: "Helps with filtering text",
	Long:  `Floppa - big russian cat.`,
	// Validation of positional arguments
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Pattern expected")
		}
		fConfig.Pattern = regexp.MustCompile(args[0])

		if len(args) == 2 {
			fConfig.FromFile = true
			fConfig.Path = args[1]
		}

		return nil
	},
	// Run logic
	Run: func(cmd *cobra.Command, args []string) {
		filter.Run(fConfig)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&fConfig.After, "after", "A", 0, "specify number of lines after match")
	rootCmd.PersistentFlags().IntVarP(&fConfig.Before, "before", "B", 0, "specify number of lines before match")
	rootCmd.PersistentFlags().IntVarP(&fConfig.Context, "context", "C", 0, "specify number before and after match")

	rootCmd.PersistentFlags().BoolVarP(&fConfig.Count, "count", "c", false, "count number of strings")
	rootCmd.PersistentFlags().BoolVarP(&fConfig.IgnoreRegister, "ignore-case", "i", false, "ignore register")
	rootCmd.PersistentFlags().BoolVarP(&fConfig.Invert, "invert", "v", false, "exclude matches")
	rootCmd.PersistentFlags().BoolVarP(&fConfig.Fixed, "fixed", "F", false, "an exact match, not a pattern")
	rootCmd.PersistentFlags().BoolVarP(&fConfig.Number, "line-num", "n", false, "show match line number")
}
