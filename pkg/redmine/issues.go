package redmine

import (
	"context"
	"fmt"
	"net/http"
)

type GetIssuesOpts struct {
}

type ListIssuesOpts struct {
}

type Issues interface {
	Get(ctx context.Context, id int, opts GetIssuesOpts) (*Issue, error)
	List(ctx context.Context, opts ListIssuesOpts) ([]Issue, error)
}

type issuesHandler struct {
	client Client
}

func newIssuesHandler(client Client) *issuesHandler {
	return &issuesHandler{client}
}

func (ih *issuesHandler) Get(ctx context.Context, id int, opts GetIssuesOpts) (*Issue, error) {
	uri := fmt.Sprintf("/issues/%d.json", id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ih.client.doRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrIssueNotFound
	}

	type getIssueResponse struct {
		Issue Issue `json:"issue"`
	}

	respBody, err := readBody[getIssueResponse](resp)
	if err != nil {
		return nil, err
	}

	return &respBody.Issue, nil
}

func (ih *issuesHandler) List(ctx context.Context, opts ListIssuesOpts) ([]Issue, error) {
	uri := "/projects.json"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ih.client.doRequest(req)
	if err != nil {
		return nil, err
	}

	type listIssuesResponse struct {
		Issues     []Issue `json:"issues"`
		TotalCount int     `json:"total_count"`
		Offset     int     `json:"offset"`
		Limit      int     `json:"limit"`
	}

	respBody, err := readBody[listIssuesResponse](resp)
	if err != nil {
		return nil, err
	}

	return respBody.Issues, nil
}
