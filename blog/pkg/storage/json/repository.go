package json

import (
	"encoding/json"
	"log"
	"path"
	"runtime"
	"time"

	"blog/pkg/adding"
	"blog/pkg/listing"
	"blog/pkg/storage"

	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines name of directory where files are stored
	dir = "/data/"

	// CollectionPost identifier for the JSON collection of posts
	CollectionPost = "posts"
)

// Storage stores post data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddPost saves the given post to the repository
func (s *Storage) AddPost(p adding.Post) error {
	id, err := storage.GenID("post")
	if err != nil {
		log.Fatal(err)
	}

	newP := Post{
		ID:      id,
		Created: time.Now(),
		Body:    p.Body,
	}

	if err := s.db.Write(CollectionPost, newP.ID, newP); err != nil {
		return err
	}
	return nil
}

// GetPost returns post with the specified ID
func (s *Storage) GetPost(id string) (listing.Post, error) {
	var p Post
	var post listing.Post

	if err := s.db.Read(CollectionPost, id, &p); err != nil {
		// err handling omitted for simplicity
		return post, listing.ErrNotFound
	}

	post.ID = p.ID
	post.Body = p.Body
	post.Created = p.Created

	return post, nil
}

// GetAllPosts returns a slice of all posts
func (s *Storage) GetAllPosts() []listing.Post {
	list := []listing.Post{}

	records, err := s.db.ReadAll(CollectionPost)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var p Post
		var post listing.Post

		if err := json.Unmarshal([]byte(r), &p); err != nil {
			// err handling omitted for simplicity
			return list
		}

		post.ID = p.ID
		post.Body = p.Body
		post.Created = p.Created

		list = append(list, post)
	}

	return list
}
