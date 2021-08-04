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

// Parser parses the given payload and extracts information out of it.
type Parser struct {
	Payload Payload
}

// NewParser sets up the parser with the payload.
func NewParser(payload string) (Parser, error) {
	var p Payload
	if err := json.Unmarshal([]byte(payload), &p); err != nil {
		return Parser{}, err
	}
	return Parser{
		Payload: p,
	}, nil
}

// RepoName gets the repo name from a payload.
func (p Parser) RepoName() string {
	return p.Payload.Repo.FullName
}

// GitURL gets the git url from a payload.
func (p Parser) GitURL() string {
	return p.Payload.Repo.GitURL
}

// SSHURL gets the ssh url from a payload.
func (p Parser) SSHURL() string {
	return p.Payload.Repo.SSHURL
}

// HTMLURL gets the html url from a payload.
func (p Parser) HTMLURL() string {
	return p.Payload.Repo.HTMLURL
}
