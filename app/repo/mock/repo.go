package mock

import "github.com/DarkSoul94/golang-template/app"

// RepoMock ...
type repoMock struct {
}

// NewRepo ...
func NewMockRepo() app.Repository {
	return &repoMock{}
}

// Close ...
func (r *repoMock) Close() error {
	return nil
}
