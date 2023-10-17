package filter

import (
	"bufio"
	"errors"
	"github.com/spf13/cobra"
	"io"
	"regexp"
	"strings"
)

type Config struct {
	Pattern        *regexp.Regexp
	Reader         *bufio.Reader
	Writer         io.Writer
	After          int64
	Before         int64
	Context        int64
	Number         bool
	Count          bool
	IgnoreRegister bool
	Invert         bool
	Fixed          bool
	FromFile       bool
	Path           string
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func getInt64(cmd *cobra.Command, name string) int64 {
	val, err := cmd.PersistentFlags().GetInt64(name)
	checkError(err)
	return val
}

func getBool(cmd *cobra.Command, name string) bool {
	val, err := cmd.PersistentFlags().GetBool(name)
	checkError(err)
	return val
}

func FromCMD(cmd *cobra.Command, args []string) *Config {

	result := &Config{
		After:   getInt64(cmd, "after"),
		Before:  getInt64(cmd, "before"),
		Context: getInt64(cmd, "context"),

		Number:         getBool(cmd, "line-num"),
		Count:          getBool(cmd, "count"),
		IgnoreRegister: getBool(cmd, "ignore-case"),
		Invert:         getBool(cmd, "invert"),
		Fixed:          getBool(cmd, "fixed"),
	}

	if len(args) < 1 {
		checkError(errors.New("Pattern expected"))
	}
	if result.IgnoreRegister {
		result.Pattern = regexp.MustCompile(strings.ToLower(args[0]))
	} else {
		result.Pattern = regexp.MustCompile(args[0])
	}

	if len(args) == 2 {
		result.FromFile = true
		result.Path = args[1]
	}

	result.After = max(result.After, result.Context)
	result.Before = max(result.Before, result.Context)

	return result
}
