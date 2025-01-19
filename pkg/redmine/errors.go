package redmine

import "errors"

var (
	ErrProjectNotFound = errors.New("project not found")
	ErrIssueNotFound   = errors.New("issue not found")
)
