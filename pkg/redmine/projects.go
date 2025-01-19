package redmine

import (
	"context"
	"fmt"
	"net/http"
)

type GetProjectOpts struct {
}

type ListProjectOpts struct {
}

type Projects interface {
	Get(ctx context.Context, identifier string, opts GetProjectOpts) (*Project, error)
	List(ctx context.Context, opts ListProjectOpts) ([]Project, error)
}

type projectsHandler struct {
	client Client
}

func newProjectsHandler(client Client) *projectsHandler {
	return &projectsHandler{client}
}

func (ph *projectsHandler) Get(ctx context.Context, identifier string, opts GetProjectOpts) (*Project, error) {
	uri := fmt.Sprintf("/projects/%s.json", identifier)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	resp, err := ph.client.doRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrProjectNotFound
	}

	type getProjectResponse struct {
		Project Project `json:"project"`
	}

	respBody, err := readBody[getProjectResponse](resp)
	if err != nil {
		return nil, err
	}

	return &respBody.Project, nil
}

func (ph *projectsHandler) List(ctx context.Context, opts ListProjectOpts) ([]Project, error) {
	uri := "/projects.json"

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	// TODO: Fetch all projects if paginated
	resp, err := ph.client.doRequest(req)
	if err != nil {
		return nil, err
	}

	type listProjectsResponse struct {
		Projects   []Project `json:"projects"`
		TotalCount int       `json:"total_count"`
		Offset     int       `json:"offset"`
		Limit      int       `json:"limit"`
	}

	respBody, err := readBody[listProjectsResponse](resp)
	if err != nil {
		return nil, err
	}

	return respBody.Projects, nil
}
