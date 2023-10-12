package filter

import (
	"bufio"
	"io"
	"log"
	"os"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readLines(reader io.Reader) []string {
	scanner := bufio.NewScanner(reader)

	result := make([]string, 0)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func doFiltering(input []string, config *Config) error {
	for i := 0; i < len(input); i++ {

	}
	return nil
}

func Run(config *Config) {
	if config.FromFile {
		file, err := os.Open(config.Path)
		check(err)
		config.Reader = bufio.NewReader(file)
		defer file.Close()
	} else {
		config.Reader = bufio.NewReader(os.Stdin)
	}

	lines := readLines(config.Reader)
	check(doFiltering(lines, config))
}
