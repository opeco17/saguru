package main

import "time"

type GithubUser struct {
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

type GithubLabel struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
}

type GithubReactions struct {
	URL        string `json:"url"`
	TotalCount int    `json:"total_count"`
	Laugh      int    `json:"laugh"`
	Hooray     int    `json:"hooray"`
	Confused   int    `json:"confused"`
	Heart      int    `json:"heart"`
	Rocket     int    `json:"rocket"`
	Eyes       int    `json:"eyes"`
}

type GithubItem struct {
	URL                   string `json:"url"`
	RepositoryURL         string `json:"repository_url"`
	LabelsURL             string `json:"labels_url"`
	CommentsURL           string `json:"comments_url"`
	EventsURL             string `json:"events_url"`
	HTMLURL               string `json:"html_url"`
	ID                    int    `json:"id"`
	NodeID                string `json:"node_id"`
	Number                int    `json:"number"`
	Title                 string `json:"title"`
	GithubUser            `json:"user"`
	GithubLabels          []GithubLabel   `json:"labels"`
	State                 string          `json:"state"`
	Locked                bool            `json:"locked"`
	Assignee              interface{}     `json:"assignee"`
	Assignees             []interface{}   `json:"assignees"`
	Milestone             interface{}     `json:"milestone"`
	Comments              int             `json:"comments"`
	CreatedAt             time.Time       `json:"created_at"`
	UpdatedAt             time.Time       `json:"updated_at"`
	ClosedAt              interface{}     `json:"closed_at"`
	AuthorAssociation     string          `json:"author_association"`
	ActiveLockReason      interface{}     `json:"active_lock_reason"`
	Body                  string          `json:"body"`
	GithubReactions       GithubReactions `json:"reactions"`
	TimelineURL           string          `json:"timeline_url"`
	PerformedViaGithubApp interface{}     `json:"performed_via_github_app"`
	Score                 float64         `json:"score"`
}

type GithubIssue struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	GithubItems       []GithubItem `json:"items"`
}
