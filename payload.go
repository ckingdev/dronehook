package dronehook

import (
	"encoding/json"
	"fmt"
)

type RepositoryType struct {
	Remote       string `json:"remote"`
	Host         string `json:"host"`
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	URL          string `json:"url"`
	CloneURL     string `json:"clone_url"`
	GitURL       string `json:"git_url"`
	SSHURL       string `json:"ssh_url"`
	Active       bool   `json:"active"`
	Private      bool   `json:"private"`
	Privileged   bool   `json:"privileged"`
	PostCommits  bool   `json:"post_commits"`
	PullRequests bool   `json:"pull_requests"`
	Timeout      int    `json:"timeout"`
	CreatedAt    int    `json:"created_at"`
	UpdatedAt    int    `json:"updated_at"`
}

type CommitType struct {
	ID          int    `json:"id"`
	Status      string `json:"status"`
	StartedAt   int    `json:"started_at"`
	FinishedAt  int    `json:"finished_at"`
	Duration    int    `json:"duration"`
	SHA         string `json:"sha"`
	Branch      string `json:"branch"`
	PullRequest string `json:"pull_request"`
	Author      string `json:"author"`
	Gravatar    string `json:"gravatar"`
	Timestamp   string `json:"timestamp"`
	Message     string `json:"message"`
	CreatedAt   int    `json:"created_at"`
	UpdatedAt   int    `json:"updated_at"`
}

type Payload struct {
	FromURL    string         `json:"from_url"`
	Commit     CommitType     `json:"commit"`
	Repository RepositoryType `json:"repository"`
}

func makePayload(raw []byte) (*Payload, error) {
	var p Payload
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (p *Payload) String() string {
	var emoji string
	if p.Commit.Status == "Success" {
		emoji = ":white_check_mark:"
	} else {
		emoji = ":x:"
	}
	return fmt.Sprintf("%s [ drone.io | %s | %s ] %s | %s",
		emoji,
		p.Repository.Name,
		p.Commit.Branch,
		p.Commit.Message,
		p.Commit.Status,
	)
}
