package github

import (
	"encoding/json"
	"fmt"
)

// ExtractRepoName gets the repo name from a payload.
func ExtractRepoName(payload string) (string, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(payload), &m); err != nil {
		return "", err
	}
	repository, ok := m["repository"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("repository was not found in payload")
	}
	fullName, ok := repository["full_name"].(string)
	if !ok {
		return "", fmt.Errorf("full_name was not found in repository: %v", repository)
	}

	return fullName, nil
}
