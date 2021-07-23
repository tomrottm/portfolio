package listing

import "errors"

// ErrNotFound is used when a post could not be found.
var ErrNotFound = errors.New("Post not found")

// Repository provides access to the post storage.
type Repository interface {
	// GetPost returns the post with given ID.
	GetPost(string) (Post, error)
	// GetAllPosts returns all posts saved in storage.
	GetAllPosts() []Post
}

// Service provides post listing operations (see below).
type Service interface {
	GetPost(string) (Post, error)
	GetPosts() []Post
}

type service struct {
	r Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// GetPost returns a post with specified id
func (s *service) GetPost(id string) (Post, error) {
	return s.r.GetPost(id)
}

// GetPosts returns all posts
func (s *service) GetPosts() []Post {
	return s.r.GetAllPosts()
}
