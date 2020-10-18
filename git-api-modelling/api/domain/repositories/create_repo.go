package repositories

import "strings"

type CreateRepoRequest struct {
	Name        string `json:"string"`
	Description string `json:"string"`
}

func (r *CreateRepoRequest) Validate() errors.ApiError {
	r.Name = strings.TrimSpace(r.Name)
	if r.Name == "" {
		return errors.NewBadRequestError("invalid repository name")
	}
	return nil
}

type CreateRepoResponse struct {
	ID    int64  `json:"id"`
	Owner string `json:"owner"`
	Name  string `json:"string"`
}

type CreateReposResponse struct {
	StatusCode int                        `json:"status"`
	Results    []CreateRepositoriesResult `json:"results"`
}

type CreateRepositoriesResult struct {
	Response *CreateRepoResponse `json:"repo"`
	Error    errors.ApiError     `json:"error"`
}
