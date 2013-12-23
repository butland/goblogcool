package post

import (
	"time"
	"io/ioutil"
	"appengine"
	"appengine/blobstore"
	"mdbook/blackfriday"
	"errors"
)

var (
	ErrDuplicatePath = errors.New("该PATH地址已经存在")
)

// 页面
// 主版本号位根据Path自动生成，次版本号为页面版本，最新版本号在Path中记录
type Post struct {
	Id		int64
	Title 	string
	Content	string
	Author 	string
	Date 	time.Time
	Version	int64
}
// 页面最新版本，提高查询性能和管理时列表所用
// 该对象仅示意，不使用，实际使用对象依然为Post，只是存储名为Snapshot
type Snapshot struct {
	Post
}

func (page *Post) ParseHTML() *Post{
	html := blackfriday.MarkdownBasic([]byte(page.Content))
	page.Content = string(html)
	lc := time.FixedZone("UTC", 28800)
	page.Date = page.Date.In(lc)
	return page
}

func (page *Post) saveBlob(ctx appengine.Context) error {
	writer, err := blobstore.Create(ctx, "text/plain")
	if err != nil {
		return err
	}
	_, err = writer.Write([]byte(page.Content))
	if err != nil {
		return err
	}
	err = writer.Close()
	if err != nil {
		return err
	}
	key, _ := writer.Key()
	page.Content = string(key)
	return nil
}

func (page *Post) loadBlob(ctx appengine.Context) error {
	key := appengine.BlobKey(page.Content)
	reader := blobstore.NewReader(ctx, key)

	bits, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	page.Content = string(bits)
	return nil
}
