package mdbook

import (
	"net/http"
	"mdbook/gorilla/mux"
	"mdbook/views"
)

func init() {
	router := mux.NewRouter()

	router.HandleFunc("/_sys/admin/page/new", views.XhrNewPostHandler)
	router.HandleFunc("/_sys/admin/page/get",  views.XhrGetPostHandler)
	router.HandleFunc("/_sys/admin/pages/get", views.XhrGetPostsHandler);
	router.HandleFunc("/_sys/admin/paths/get", views.XhrGetPathsHandler);
	router.HandleFunc("/_sys/admin/paths/max", views.XhrGetMaxHandler);

	router.HandleFunc("/_sys/admin/help",  views.HelpHandler )
	router.HandleFunc("/_sys/admin/page",  views.PostEditHandler )
	router.HandleFunc("/_sys/admin/upload",  views.UploadHandler )
	router.HandleFunc("/_sys/admin",  views.PostListHandler )

	router.HandleFunc("/_sys/go/show/{id}",  views.GoShowPostHandler )
	router.HandleFunc("/_sys/go/edit/{id}",  views.GoEditPostHandler )

	router.HandleFunc("/_sys/{path:.*}",  views.ReservePageHandler )

	router.HandleFunc("/file/{id}",  views.DownloadHandler )
	router.HandleFunc("/file/{id}/{path:.*}",  views.ReservePageHandler )
	router.HandleFunc("/{path:.*}",  views.PostHandler)

	http.Handle("/", router)
}
