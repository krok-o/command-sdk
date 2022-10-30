package pkg

import (
	"fmt"
	"strings"
)

// Output formats some values as output to another command.
// All output should be put on error.
func Output(m map[string]string) string {
	var output []string
	for k, v := range m {
		output = append(output, fmt.Sprintf("%s=%s", k, v))
	}

	return strings.Join(output, "\n")
}
