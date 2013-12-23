package views

import (
	"strconv"
	"net/http"
	"appengine"
	"mdbook/gorilla/mux"
	"mdbook/model/post"
)

func GoShowPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	path, err := post.GetPathById(id, ctx)
	if err != nil {
		renderError(w, r, err)
	}
	http.Redirect(w, r, path.Path, 301)
}

func GoEditPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	path, err := post.GetPathById(id, ctx)
	if err != nil {
		renderError(w, r, err)
	}
	http.Redirect(w, r, "/_sys/admin/page?path=" + path.Path, 301)
}
