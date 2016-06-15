package webpage

import (
	"gin_miaomiaola/dao"
	"gin_miaomiaola/mkd"
	"html/template"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
)

//HomePage 首页
func HomePage(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)
	log.Info("展现首页处理开始")

	db := c.MustGet("db").(*mgo.Database)

	p, err := dao.GetAllTopic(db, log)
	if err != nil {
		log.Error("获取文章出错", err)
		c.HTML(404, "home.tmpl", p)
	}
	c.HTML(200, "home.tmpl", p)
}

//PostPage 详情
func PostPage(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)
	log.Info("展现内容处理开始")

	db := c.MustGet("db").(*mgo.Database)
	url := c.Param("urlString")

	if url == "" {
		log.Error("参数为空")
		c.HTML(404, "404.tmpl", struct{}{})
		return
	}

	p, err := dao.GetTopic(url, db, log)
	if err != nil {
		log.Error("获取文章出错", err)
		c.HTML(404, "home.tmpl", p)
	}
	rtn := new(Post)
	rtn.HTML = template.HTML(mkd.MarkdownToHTML(p.Desc, log))
	rtn.Post = *p
	log.Debug("Markdown 后的数据", "DESC", rtn.HTML)
	c.HTML(200, "postDetail.tmpl", rtn)
}

//AboutMe 详情
func AboutMe(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)
	log.Info("展现内容处理开始")

	db := c.MustGet("db").(*mgo.Database)
	url := c.Param("urlString")

	if url == "" {
		log.Error("参数为空")
		c.HTML(404, "404.tmpl", struct{}{})
		return
	}

	p, err := dao.GetTopic(url, db, log)
	if err != nil {
		log.Error("获取文章出错", err)
		c.HTML(404, "home.tmpl", p)
	}
	rtn := new(Post)
	rtn.HTML = template.HTML(mkd.MarkdownToHTML(p.Desc, log))
	rtn.Post = *p
	log.Debug("Markdown 后的数据", "DESC", rtn.HTML)
	c.HTML(200, "about.tmpl", rtn)
}
