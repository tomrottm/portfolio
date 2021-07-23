package rest

import (
	"encoding/json"
	"net/http"

	"blog/pkg/adding"
	"blog/pkg/listing"

	"github.com/julienschmidt/httprouter"
)

func Handler(a adding.Service, l listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/posts", getPosts(l))
	router.GET("/posts/:id", getPost(l))

	return router
}

func getPosts(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")

		list := s.GetPosts()
		json.NewEncoder(w).Encode(list)
	}
}

func getPost(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		post, err := s.GetPost(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The post you request does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(post)
	}
}
