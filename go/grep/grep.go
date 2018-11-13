package grep

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Search(pattern string, flags []string, files []string) []string {
	result := []string{}
	argsMap := map[string]bool{}
	isMultiFile := len(files) > 1
	for _, arg := range flags {
		argsMap[arg] = true
	}
	for _, fname := range files {
		file, err := os.Open(fname)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for i := 1; scanner.Scan(); i++ {
			line := scanner.Text()
			matched := line
			if argsMap["-i"] {
				matched = strings.ToLower(line)
				pattern = strings.ToLower(pattern)
			}
			match := strings.Contains(matched, pattern)
			if argsMap["-x"] {
				match = matched == pattern
			}
			if (argsMap["-v"] && !match) || (!argsMap["-v"] && match) {
				if argsMap["-n"] {
					line = fmt.Sprintf("%d:%s", i, line)
				}
				if isMultiFile {
					line = fmt.Sprintf("%s:%s", fname, line)
				}
				if argsMap["-l"] {
					result = append(result, fname)
					break
				}
				result = append(result, line)
			}
		}
	}
	return result
}
