package views

import (
	"bytes"
	"strconv"
	"net/http"
	"net/url"
	"encoding/json"
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"mdbook/model/post"
)

type content struct {
	Title		string
	ThemePath	string
	Header		string
	Content		string
	Footer		string
}

type header struct {
	CanBack		bool
	Logined		bool
	IsAdmin		bool
	User		string
	LoginUrl	string
	LogoutUrl	string
}

type postExt struct {
	*post.Post
	Path	string
}

// 公共页面处理
func parseTemplate(file string, data interface{}) (out []byte, error error) {
	var buf bytes.Buffer
	err := tmpls.ExecuteTemplate(&buf, file, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func parseHeader(r *http.Request, file string) ([]byte, error){
	c := appengine.NewContext(r)
	h := new(header)

	if r.URL.Path != "/" {
		ref := r.Header.Get("Referer")
		if ref != "" {
			uref, _ := url.Parse(ref)
			if uref.Host == r.Host {
				h.CanBack = true
			}
		}
	}

	if u := user.Current(c); u != nil {
		h.Logined = true
		h.User = u.String()
		if u.Admin {
			h.IsAdmin = true
		}
	}

	h.LoginUrl, _ = user.LoginURL(c, r.URL.String())
	h.LogoutUrl, _ = user.LogoutURL(c, r.URL.String())

	b, err := parseTemplate(file, h)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func renderView(r *http.Request, file string, title string, data interface{}) ([]byte, error) {

	hbyte, err := parseHeader(r, "header.tpl")
	if err != nil {
		return nil, err
	}

	footer, err := parseTemplate("footer.tpl", nil)
	if err != nil {
		return nil, err
	}
	page, err := parseTemplate(file, data)
	if err != nil {
		return nil, err
	}
	base, err := parseTemplate("base.tpl", content{Title: title, ThemePath: themePath, Header: string(hbyte), Content: string(page), Footer: string(footer)})
	if err != nil {
		return nil, err
	}
	return base, nil
}

func renderJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, _ := json.Marshal(v)
	w.Write(b)
}

func render404(w http.ResponseWriter, r *http.Request, path string) {
	w.WriteHeader(http.StatusNotFound)
	view, _ := renderView(r, "404.tpl", "404", path)
	w.Write(view)
}

func renderError(w http.ResponseWriter, r *http.Request, err error) {
	c := appengine.NewContext(r)
	w.WriteHeader(http.StatusInternalServerError)
	c.Errorf("%v", err)
	view, _ := renderView(r, "500.tpl", "500", nil)
	w.Write(view)
}

// 浏览网站
func ReservePageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	view, _ := renderView(r, "reserve.tpl", "404", nil)
	w.Write(view)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	ver, _ := strconv.ParseInt(query.Get("ver"), 10, 64)
	raw := query.Get("raw")
	ctx := appengine.NewContext(r)
	p, err := post.Get(r.URL.Path, ver, ctx)
	if err != nil {
	    if err == datastore.ErrNoSuchEntity {
			render404(w, r, r.URL.Path)
			return
		}
		renderError(w, r, err)
		return
	}
	if raw == "yes" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte(p.Content))
		return
	}

	extp := postExt{p.ParseHTML(), r.URL.Path}
	ctx.Infof("%v", extp)
	view, err := renderView(r, "post.tpl", extp.Title, extp)
	if err != nil {
		ctx.Errorf("%v", err)
		renderError(w, r, err)
		return
	}
	w.Write(view)
}

////////////////
/*
// Shows the welcome screen. This might actually be the blog itself.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	page := cache.getCachedPage(r.URL.RequestURI())
	if page == nil {
		post := getNewestBlogPost()
		page = getPage("templates/blog.html", post)
		cache.putCachedPage(r.URL.RequestURI(), page)
	}
	//fmt.Fprintf(w, page)
	w.Write(page)
}
*/
