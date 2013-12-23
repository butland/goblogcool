package views

import (
	"io/ioutil"
	"encoding/json"
	"text/template"
)

type Config struct {
	Theme	string
}

func init(){
	bts, _ := ioutil.ReadFile("config.json")
	cfg := new(Config)
	_ = json.Unmarshal(bts, cfg);

	theme = cfg.Theme
	themePath = "themes/" + theme
	tmpls = template.Must(template.ParseFiles(
		// 用户模板
		themePath + "/base.tpl",
		themePath + "/header.tpl",
		themePath + "/footer.tpl",
		themePath + "/post.tpl",
		themePath + "/reserve.tpl",
		themePath + "/404.tpl",
		themePath + "/500.tpl",
	))
}

var (
	// 皮肤名称
	theme =  "base"
	themePath = "themes/" + theme
	tmpls = template.New("tmpls")
)
