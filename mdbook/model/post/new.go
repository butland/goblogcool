package post

import (
	"time"
	"strings"
	"appengine"
	"appengine/datastore"
	"appengine/user"
)

func New(path string, content string, title string, ctx appengine.Context) (*Path, error){
	// 创建Path
	pa := &Path{ Path:path }
	err := pa.New(ctx)
	if err != nil {
		return nil, err
	}

	post := &Post{Id: pa.Id, Version:pa.Latest}
	if usr := user.Current(ctx); usr != nil {
		post.Author = usr.String()
	}
	// 删除多余的空格
	post.Content = strings.Trim(content, " ")
	post.Title = title
	post.Date =	time.Now()

	ppostKey := datastore.NewKey(ctx, "Post", "", pa.Id, nil)
	postKey := datastore.NewKey(ctx, "Post", "", pa.Latest, ppostKey)

	// 保存内容到blob中
	if err := post.saveBlob(ctx) ; err != nil {
		// 保存失败后恢复版本信息
		pa.Decrease(ctx)
		return nil, err
	}
	// 保存内容到db中
	if  _, err := datastore.Put(ctx, postKey, post); err != nil {
		// 保存失败后恢复版本信息
		pa.Decrease(ctx)
		return nil, err
	}
	// 更新该页快照
	snapKey := datastore.NewKey(ctx, "Snapshot", "", pa.Id, nil)
	if  _, err := datastore.Put(ctx, snapKey, post); err != nil {
		return nil, err
	}
	return pa, nil
}

// 修改 Post Path地址
func RedirectPath(from string, to string, ctx appengine.Context)(*Path, error){
	// 判断新URL是否已经被使用
	tokey := datastore.NewKey(ctx, "Path", to, 0, nil)
	tp := new(Path)
	if err := datastore.Get(ctx, tokey, tp); err == nil {
		return nil, ErrDuplicatePath
	} else if err != datastore.ErrNoSuchEntity {
		return nil, err
	}

	fromkey := datastore.NewKey(ctx, "Path", from, 0, nil)
	fp := new(Path)
	if err := datastore.Get(ctx, fromkey, fp); err != nil {
		return nil, err
	}

	tp = fp
	tp.Path = to
	if _, err := datastore.Put(ctx, tokey, tp); err != nil {
		return nil, err
	}
	fp.NewPath = to
	if _, err := datastore.Put(ctx, fromkey, fp); err != nil {
		return nil, err
	}
	return tp, nil
}
