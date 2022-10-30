package pkg

import (
	"fmt"
	"os"
	"strings"
)

var (
	BeginOutputFormat = "----- BEGIN OUTPUT -----"
	EndOutputFormat   = "----- END OUTPUT -----"
)

// Output formats some values as output to another command.
// All output should be put on error.
func Output(m map[string]string) error {
	output := []string{
		BeginOutputFormat,
	}
	for k, v := range m {
		output = append(output, fmt.Sprintf("%s=%s", k, v))
	}
	output = append(output, EndOutputFormat)
	if _, err := fmt.Fprint(os.Stderr, strings.Join(output, "\n")); err != nil {
		return fmt.Errorf("failed to write to stderr: %w", err)
	}
	return nil
}
