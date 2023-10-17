package filter

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func normalize(src string, config *Config) string {
	if config.IgnoreRegister {
		return strings.ToLower(src)
	}
	return src
}

func checkMatch(src string, config *Config) (result bool) {
	if config.Fixed {
		indexes := config.Pattern.FindString(normalize(src, config))
		result = len(indexes) != 0
	} else {
		result = config.Pattern.MatchString(normalize(src, config))
	}

	if config.Invert {
		result = !result
	}
	return
}

func printLine(line string, num int64, config *Config) {
	if config.Number {
		fmt.Printf("%d) %v\n", num, line)
	} else {
		fmt.Println(line)
	}
}

func printResults(window map[int64]string, config *Config, left, right int64) {

	for i := left; i < right; i++ {
		printLine(window[i], i, config)
	}

	if config.Before != 0 || config.After != 0 {
		fmt.Println("END:")
	}
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

	windowCap := config.After + config.Before
	window := make(map[int64]string, windowCap)
	var left int64 = 1
	var right int64 = 1
	var curr int64 = 0
	var matched int64 = 0

	scanner := bufio.NewScanner(config.Reader)

	for curr != right {
		curr++
		for right < (curr + config.After + 1) {
			if ok := scanner.Scan(); ok {
				line := scanner.Text()
				window[right] = line
				right++
			} else {
				break
			}
		}

		for ; left <= curr-config.Before-1; left++ {
			delete(window, left)
		}

		if line, ok := window[curr]; ok {
			if checkMatch(line, config) {
				printResults(window, config, left, right)
				matched++
			}
		}

	}

	if config.Count {
		fmt.Printf("Lines: %d\n", matched)
	}
}
