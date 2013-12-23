package post

import (
	"appengine"
	"appengine/datastore"
)

// 根据页面版本获取内容
// 		path 页面URL
//		version 页面版本，当 version 小于等于0 表示未设置版本号，那么则取最新的版本
func Get(path string, version int64, ctx appengine.Context) (*Post, error){
	p, err := getPath(path, ctx)
	if err != nil {
		return nil, err
	}
	if version <= 0 {
		post, err := getSnapshot(p.Id, ctx)
		return post, err
	}

	post, err := getHistory(p.Id, version, ctx)
	return post, err
}

func getPath(path string, ctx appengine.Context) (*Path, error){
	p := new(Path)
	pkey := datastore.NewKey(ctx, "Path", path, 0, nil);

	if err := datastore.Get(ctx, pkey, p); err != nil {
		return nil, err
	}
	return p, nil
}

func getHistory(id int64, version int64, ctx appengine.Context) (*Post, error) {
	/*var post *Post
	// 页面主版本由 PATH 决定
	pkey := datastore.NewKey(ctx, "Post", "", id, nil)
	// 各子版本有版本号决定
	key := datastore.NewKey(ctx, "Post", "", version, pkey);
	if err := datastore.Get(ctx, key, post); err != nil {
		return nil, err
	}

	if err := post.loadBlob(ctx); err != nil {
		return nil, err
	}
	return post, nil*/

	var ps []*Post
	query := datastore.NewQuery("Post").Filter("Id=", id).Filter("Version=",version).Limit(1)
	_, err := query.GetAll(ctx, &ps)
	if err != nil {
		return nil, err
	} else if len(ps) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}

	post := ps[0]
	if err := post.loadBlob(ctx); err != nil {
		return nil, err
	}
	return post, nil
}

func getSnapshot(id int64, ctx appengine.Context)(*Post, error){
	key := datastore.NewKey(ctx, "Snapshot", "", id, nil);
	snap := new(Post)
	if err := datastore.Get(ctx, key, snap); err != nil {
		return nil, err
	}
	if err := snap.loadBlob(ctx); err != nil {
		return nil, err
	}
	return snap, nil
}

func GetPaths(start int, to int, ctx appengine.Context) ([]*Path, error){
	if start > to {
		tmp := start
		start = to
		to = tmp
	}
	var ps []*Path
	query := datastore.NewQuery("Path").Order("-Id").Filter("Id>=", start).Filter("Id<=",to).Filter("NewPath=","")
	_, err := query.GetAll(ctx, &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func GetPosts(offset int, limit int, order string, ctx appengine.Context) ([]*Post, error){
	var ps []*Post
	query := datastore.NewQuery("Snapshot").Order(order).Offset(offset).Limit(limit)
	_, err := query.GetAll(ctx, &ps)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func GetPathById(id int64, ctx appengine.Context)(*Path, error) {
	var ps []*Path
	query := datastore.NewQuery("Path").Filter("Id=", id).Filter("NewPath=", "").Limit(1)
	_, err := query.GetAll(ctx, &ps)
	if err != nil {
		return nil, err
	}
	if len(ps) == 0 {
		return nil, datastore.ErrNoSuchEntity
	}
	return ps[0], nil
}

