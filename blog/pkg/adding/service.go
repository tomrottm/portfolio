package adding

import (
	"blog/pkg/listing"
)

// Service provides post adding operations.
type Service interface {
	AddPost(...Post) error
	AddDefaultPosts([]Post)
}

// Repository provides access to post repository.
type Repository interface {
	// AddPost saves a given post to the repository.
	AddPost(Post) error
	// GetAllPosts returns all posts saved in storage.
	GetAllPosts() []listing.Post
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) AddPost(p ...Post) error {
	// implement: check for duplicates

	for _, post := range p {
		_ = s.r.AddPost(post) // error handling omitted for simplicity
	}

	return nil
}

// AddDefaultPosts adds some sample posts to the database
func (s *service) AddDefaultPosts(p []Post) {

	// any validation can be done here

	for _, pp := range p {
		_ = s.r.AddPost(pp) // error handling omitted for simplicity
	}
}
