package mock

import "github.com/DarkSoul94/golang-template/app"

// RepoMock ...
type repoMock struct {
}

// NewRepo ...
func NewRepo() app.Repository {
	return &repoMock{}
}

// Close ...
func (r *repoMock) Close() error {
	return nil
}
