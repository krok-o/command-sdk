package github

import (
	"encoding/json"
)

// Repository contains information about the repository. All we care about
// here are the possible urls for identification.
type Repository struct {
	GitURL   string `json:"git_url"`
	SSHURL   string `json:"ssh_url"`
	HTMLURL  string `json:"html_url"`
	FullName string `json:"full_name"`
}

// Payload contains information about the event like, user, commit id and so on.
// All we care about for the sake of identification is the repository.
type Payload struct {
	Repo Repository `json:"repository"`
}

// ExtractRepoName gets the repo name from a payload.
func ExtractRepoName(payload string) (string, error) {
	var p Payload
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		return "", err
	}

	return p.Repo.FullName, nil
}
