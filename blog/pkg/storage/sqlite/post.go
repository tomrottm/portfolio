package sqlite

import (
	"time"
)

// Post defines the storage form of a post
type Post struct {
	ID     int
	UserID int
	Body   string
	Time   time.Time
}
