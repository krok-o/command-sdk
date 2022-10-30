package pkg

import (
	"fmt"
	"os"
	"strings"
)

// Output formats some values as output to another command.
// All output should be put on error.
func Output(m map[string]string) error {
	var output []string
	for k, v := range m {
		output = append(output, fmt.Sprintf("%s=%s", k, v))
	}

	if _, err := fmt.Fprint(os.Stderr, strings.Join(output, "\n")); err != nil {
		return fmt.Errorf("failed to write to stderr: %w", err)
	}
	return nil
}
