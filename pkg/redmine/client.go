package redmine

import (
	"net/http"
	"net/url"
	"strings"
)

type Client interface {
	doRequest(req *http.Request) (*http.Response, error)
	Projects() Projects
	Issues() Issues
}

type APIClient struct {
	config     Config
	httpClient *http.Client
	projects   Projects
	issues     Issues
}

func NewAPIClient(config Config) *APIClient {
	client := &APIClient{
		config:     config,
		httpClient: http.DefaultClient,
	}

	client.projects = newProjectsHandler(client)
	client.issues = newIssuesHandler(client)
	return client
}

func (a *APIClient) doRequest(req *http.Request) (*http.Response, error) {
	if !strings.HasPrefix(req.URL.String(), a.config.BaseURL) {
		fullURL, err := url.Parse(a.config.BaseURL + req.URL.Path)
		if err != nil {
			return nil, err
		}
		req.URL = fullURL
	}

	req.Header.Add("X-Redmine-API-Key", a.config.APIKey)

	return a.httpClient.Do(req)
}

func (a *APIClient) Projects() Projects {
	return a.projects
}

func (a *APIClient) Issues() Issues {
	return a.issues
}
