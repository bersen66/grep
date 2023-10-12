package filter

import (
	"bufio"
	"io"
	"regexp"
)

type Config struct {
	Pattern        *regexp.Regexp
	Reader         *bufio.Reader
	Writer         io.Writer
	After          int
	Before         int
	Context        int
	Number         bool
	Count          bool
	IgnoreRegister bool
	Invert         bool
	Fixed          bool
	FromFile       bool
	Path           string
}
