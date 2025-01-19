package redmine

import "time"

type IdName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Multiple    bool   `json:"multiple"`
	Value       any    `json:"value"`
}

type Project struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	Identifier     string        `json:"identifier"`
	Description    string        `json:"description"`
	Homepage       string        `json:"homepage"`
	Status         int           `json:"status"`
	IsPublic       bool          `json:"is_public"`
	InheritMembers bool          `json:"inherit_members"`
	CreatedOn      time.Time     `json:"created_on"`
	UpdatedOn      time.Time     `json:"updated_on"`
	Parent         *IdName       `json:"parent,omitempty"`
	CustomFields   []CustomField `json:"custom_fields,omitempty"`
}

type Issue struct {
	Id             int           `json:"id"`
	Subject        string        `json:"subject"`
	Description    string        `json:"description"`
	ProjectId      int           `json:"project_id"`
	Project        *IdName       `json:"project"`
	TrackerId      int           `json:"tracker_id"`
	Tracker        *IdName       `json:"tracker"`
	ParentId       int           `json:"parent_issue_id,omitempty"`
	Parent         *int          `json:"parent"`
	StatusId       int           `json:"status_id"`
	Status         *IdName       `json:"status"`
	PriorityId     int           `json:"priority_id,omitempty"`
	Priority       *IdName       `json:"priority"`
	Author         *IdName       `json:"author"`
	FixedVersion   *IdName       `json:"fixed_version"`
	AssignedTo     *IdName       `json:"assigned_to"`
	AssignedToId   int           `json:"assigned_to_id,omitempty"`
	Category       *IdName       `json:"category"`
	CategoryId     int           `json:"category_id,omitempty"`
	Notes          string        `json:"notes"`
	StatusDate     string        `json:"status_date"`
	CreatedOn      string        `json:"created_on"`
	UpdatedOn      string        `json:"updated_on"`
	StartDate      string        `json:"start_date"`
	DueDate        string        `json:"due_date"`
	ClosedOn       string        `json:"closed_on"`
	CustomFields   []CustomField `json:"custom_fields,omitempty"`
	Uploads        []Upload      `json:"uploads"`
	DoneRatio      float32       `json:"done_ratio"`
	EstimatedHours float32       `json:"estimated_hours"`
	Journals       []Journal     `json:"journals"`
}

type Upload struct {
	Token       string `json:"token"`
	Filename    string `json:"filename"`
	ContentType string `json:"content_type"`
}

type Journal struct {
	Id        int              `json:"id"`
	User      *IdName          `json:"user"`
	Notes     string           `json:"notes"`
	CreatedOn string           `json:"created_on"`
	Details   []JournalDetails `json:"details"`
}

type JournalDetails struct {
	Property string `json:"property"`
	Name     string `json:"name"`
	OldValue string `json:"old_value"`
	NewValue string `json:"new_value"`
}
