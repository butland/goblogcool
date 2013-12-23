package views

import (
	"bytes"
	"strings"
	"strconv"
	"net/http"
	"text/template"
	"mdbook/gorilla/mux"
	"appengine"
	"appengine/blobstore"
	"appengine/datastore"
	"appengine/user"
	"mdbook/model/post"
	"mdbook/model/file"
)

var (
	admTmpls = template.Must(template.ParseFiles(
		// 管理界面模版
		"admin/base.tpl",
		"admin/header.tpl",
		"admin/footer.tpl",
		"admin/editor.tpl",
		"admin/posts.tpl",
		"admin/help.tpl",
	))
)

// 公共页面处理
func parseAdmTemplate(file string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	err := admTmpls.ExecuteTemplate(&buf, file, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseAdmHeader(r *http.Request, file string) ([]byte, error){
	c := appengine.NewContext(r)
	h := new(header)

	if u := user.Current(c); u != nil {
		h.Logined = true
		h.User = u.String()
		if u.Admin {
			h.IsAdmin = true
		}
	}

	h.LogoutUrl, _ = user.LogoutURL(c, "/")

	b, err := parseAdmTemplate(file, h)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func renderAdmView(r *http.Request, file string, title string, data interface{}) ([]byte, error) {

	hbyte, err := parseAdmHeader(r, "header.tpl")
	if err != nil {
		return nil, err
	}

	footer, err := parseAdmTemplate("footer.tpl", nil)
	if err != nil {
		return nil, err
	}
	result, err := parseAdmTemplate(file, data)
	if err != nil {
		return nil, err
	}
	base, err := parseAdmTemplate("base.tpl", content{Title: title, ThemePath: themePath, Header: string(hbyte), Content: string(result), Footer: string(footer)})
	if err != nil {
		return nil, err
	}
	return base, nil
}

// 文件处理

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	f, err := file.SaveFile(r, "file")
	if err != nil {
		c.Errorf("SAVE FILE ERROR: %v", err)
		renderJSON(w, r, &file.File{})
		return
	}

	q := datastore.NewQuery("File").Order("-Id").Limit(1)
	var ff []*file.File
	var max int64
	_, err = q.GetAll(c, &ff)
	if err != nil {
		c.Infof("GET FILE MAXID ERROR: %v", err)
		max = 1
	}
	if len(ff) == 0 {
		max = 1
	} else {
		max = ff[0].Id + 1
	}

	k := datastore.NewKey(c, "File", "", max, nil);
	f.Id = max
	if _, err := datastore.Put(c, k, f); err != nil {
		c.Errorf("SAVE FILE INFO ERROR: %v", err)
		renderJSON(w, r, &file.File{})
	}
	renderJSON(w, r, f)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	vars := mux.Vars(r)

	id, _ := strconv.ParseInt(vars["id"], 10, 64)
	k := datastore.NewKey(c, "File", "", id, nil);
	f := new(file.File)
	if err := datastore.Get(c, k, f); err != nil {
		c.Errorf("GET FILE INFO: %v",  err)
		render404(w, r, "")
	}
	if strings.Index(f.Mime, "image/") == 0 {
		w.Header().Set("Age","57162")
		w.Header().Set("Cache-Control","max-age=315360000")
		w.Header().Set("Expires","Wed, 14 Sep 2022 15:48:08 GMT")
	} else {
		w.Header().Set("Content-Disposition", `attachment; filename="` + f.Name + `"`)
	}
	blobstore.Send(w, appengine.BlobKey(f.BlobId))
}

// 管理功能
func PostEditHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	admTmpls.ExecuteTemplate(w, "editor.tpl", nil)
}

func PostListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, err := renderAdmView(r, "posts.tpl", "页面管理", nil)
	if err != nil {
		renderError(w, r, err)
	}
	w.Write(view)
}

func HelpHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	admTmpls.ExecuteTemplate(w, "help.tpl", nil)
}

func XhrGetPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	query := r.URL.Query()
	ver, _ := strconv.ParseInt(query.Get("ver"), 10, 64)
	p, err := post.Get(query.Get("path"), ver, ctx)
	if err != nil {
		renderJSON(w, r, p)
		return
	}
	renderJSON(w, r, p)
}

func XhrNewPostHandler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	page, err := post.New(r.FormValue("path"), r.FormValue("text"), r.FormValue("title"), ctx)
	if err != nil {
		renderJSON(w, r, err)
		return
	}
	renderJSON(w, r, page)
}

func XhrGetPostsHandler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	query := r.URL.Query()
	offset, _ := strconv.Atoi(query.Get("offset"))
	limit, _ := strconv.Atoi(query.Get("limit"))
	order := query.Get("order");
	if order == "" {
		order = "-Id"
	}
	ps, err := post.GetPosts(offset, limit, order, ctx)
	if err != nil {
		renderJSON(w, r, err)
		return
	}
	renderJSON(w, r, ps)
}

func XhrGetPathsHandler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	query := r.URL.Query()
	start, _ := strconv.Atoi(query.Get("start"))
	to, _ := strconv.Atoi(query.Get("to"))
	ps, err := post.GetPaths(start, to, ctx)
	if err != nil {
		renderJSON(w, r, err)
		return
	}
	renderJSON(w, r, ps)
}

func XhrGetMaxHandler(w http.ResponseWriter, r *http.Request){
	ctx := appengine.NewContext(r)
	max, err := post.GetMaxId(ctx)
	if err != nil {
		renderJSON(w, r, err)
		return
	}
	renderJSON(w, r, &post.Path{Id:max})
}

