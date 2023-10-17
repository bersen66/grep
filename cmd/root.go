package cmd

import (
	"github.com/bersen66/grep/pkg/filter"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "grep [pattern] {file} {flags}",
	Short: "Helps with filtering text",
	Long:  `Floppa - big russian cat.`,
	// Run logic
	Run: func(cmd *cobra.Command, args []string) {
		fConfig := filter.FromCMD(cmd, args)
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
	rootCmd.PersistentFlags().Int64P("after", "A", 0, "specify number of lines after match")
	rootCmd.PersistentFlags().Int64P("before", "B", 0, "specify number of lines before match")
	rootCmd.PersistentFlags().Int64P("context", "C", 0, "specify number before and after match")

	rootCmd.PersistentFlags().BoolP("count", "c", false, "count number of strings")
	rootCmd.PersistentFlags().BoolP("ignore-case", "i", false, "ignore register")
	rootCmd.PersistentFlags().BoolP("invert", "v", false, "exclude matches")
	rootCmd.PersistentFlags().BoolP("fixed", "F", false, "an exact match, not a pattern")
	rootCmd.PersistentFlags().BoolP("line-num", "n", false, "show match line number")
}
