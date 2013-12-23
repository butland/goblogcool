package file

import (
	"net/http"
	"appengine"
	"appengine/blobstore"
	"io/ioutil"
	"strings"
)

type File struct {
	Id 		int64
	BlobId 	string
	Mime 	string
	Name	string
}

func SaveFile(r *http.Request, name string ) (file *File, err error) {
	c := appengine.NewContext(r)

	fb, h, err := r.FormFile(name)
	if err != nil {
		return nil, err
	}

	fname := strings.ToLower(h.Filename)
	var mime string
	extIdx := len(fname) - 3
	if strings.Index(fname, "jpg") == extIdx {
		mime = "image/jpeg"
	} else if strings.Index(fname, "png") == extIdx {
		mime = "image/png"
	} else if strings.Index(fname, "gif") == extIdx {
		mime = "image/gif"
	} else {
		mime = "application/octet-stream"
	}
	//c.Infof("MIME INFO: %v, %v, %v", fname, mime, extIdx)
	w, err := blobstore.Create(c, mime)
		if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(fb)
	_, err = w.Write(b)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}

	k, _ := w.Key()
	f := &File{BlobId: string(k), Mime: mime, Name: fname }
	return f, nil
}
