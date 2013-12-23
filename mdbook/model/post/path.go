package post

import (
	"appengine"
	"appengine/datastore"
	"strconv"
)

// 页面Path（url）
// id 为自动生成，Path为Url，查询键值为Path
// NewPath 默认为空，当修改Path后指向新的Path，访问该Path将301重定向到新Path
// Latest 记录该页的最新版本
type Path struct {
	Id		int64
	Path	string
	NewPath	string
	Latest	int64
}

// 通过path创建Path对象
// 如果未设置 path，则根据ID自动生成
// 创建后Latest 自动加 1
func (path *Path) New(ctx appengine.Context) error{
	if e, err := path.Exists(ctx); err != nil {
		return err
	} else if e {
		path.Latest += 1
	} else {
		max, err := GetMaxId(ctx)
		if err != nil {
			return nil
		}
		if path.Path == "" {
			path.Path = "/post/" + strconv.FormatInt(max, 10)
		}
		path.Id = max
		path.Latest = 1
	}
	key := datastore.NewKey(ctx, "Path", path.Path, 0, nil)
	_, err := datastore.Put(ctx, key, path)
	return err
}

func GetMaxId(ctx appengine.Context) (int64, error){
	q := datastore.NewQuery("Path").Order("-Id").Limit(1)
	var pp []*Path
	var max int64
	_, err := q.GetAll(ctx, &pp)
	if err != nil {
		if err != datastore.ErrNoSuchEntity {
			return 0, err
		}
		max = 1
	}
	if len(pp) == 0 {
		max = 1
	} else {
		max = pp[0].Id + 1
	}
	return max, nil
}

// 恢复版本号
func (path *Path) Decrease(ctx appengine.Context) error {
	path.Latest -= 1
	key := datastore.NewKey(ctx, "Path", path.Path, 0, nil)
	_, err := datastore.Put(ctx, key, path)
	return err
}

// 判断 Path 是否存在
func (path *Path) Exists(ctx appengine.Context) (bool, error){
	if path.Path == "" {
		return false, nil
	}
	key := datastore.NewKey(ctx, "Path", path.Path, 0, nil)	
	if err := datastore.Get(ctx, key, path); err != nil {
		if err != datastore.ErrNoSuchEntity {
			return false, err
		}	
		return false, nil
	}
	return true, nil
}
