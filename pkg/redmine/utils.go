package redmine

import (
	"encoding/json"
	"io"
	"net/http"
)

// readBody reads body of the response and Unmashalls from JSON
func readBody[T any](r *http.Response) (T, error) {
	var result T

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return result, err
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
