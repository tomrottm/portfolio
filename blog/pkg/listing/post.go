package listing

import (
	"time"
)

// Post defines the storage form of a post
type Post struct {
	ID      string    `json:"id"`
	Body    string    `json:"body"`
	Created time.Time `json:"time"`
}
