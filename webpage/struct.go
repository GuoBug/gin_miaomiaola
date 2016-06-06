package webpage

import (
	"gin_miaomiaola/dao"
	"html/template"
)

//Post 文章内容
type Post struct {
	dao.Post
	HTML template.HTML
}
