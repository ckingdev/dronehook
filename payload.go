package dronehook

import (
	"encoding/json"
)

func makePayload(raw []byte) (*Payload, error) {
	var p Payload
	if err := json.Unmarshal(raw, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

type Payload struct {
	Build struct {
		Author       string `json:"author"`
		AuthorAvatar string `json:"author_avatar"`
		AuthorEmail  string `json:"author_email"`
		Branch       string `json:"branch"`
		Commit       string `json:"commit"`
		CreatedAt    int    `json:"created_at"`
		EnqueuedAt   int    `json:"enqueued_at"`
		Event        string `json:"event"`
		FinishedAt   int    `json:"finished_at"`
		LinkURL      string `json:"link_url"`
		Message      string `json:"message"`
		Number       int    `json:"number"`
		Ref          string `json:"ref"`
		Refspec      string `json:"refspec"`
		Remote       string `json:"remote"`
		StartedAt    int    `json:"started_at"`
		Status       string `json:"status"`
		Timestamp    int    `json:"timestamp"`
		Title        string `json:"title"`
	} `json:"build"`
	Repo struct {
		AllowDeploys  bool   `json:"allow_deploys"`
		AllowPr       bool   `json:"allow_pr"`
		AllowPush     bool   `json:"allow_push"`
		AllowTags     bool   `json:"allow_tags"`
		AvatarURL     string `json:"avatar_url"`
		CloneURL      string `json:"clone_url"`
		DefaultBranch string `json:"default_branch"`
		FullName      string `json:"full_name"`
		LinkURL       string `json:"link_url"`
		Name          string `json:"name"`
		Owner         string `json:"owner"`
		Private       bool   `json:"private"`
		Timeout       int    `json:"timeout"`
		Trusted       bool   `json:"trusted"`
	} `json:"repo"`
}
