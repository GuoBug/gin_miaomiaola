package webpage

import (
	"gin_miaomiaola/dao"

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

	p, err := dao.GetTopic(db, log)
	if err != nil {
		log.Error("获取文章出错", err)
		c.HTML(404, "home.tmpl", p)
	}
	c.HTML(200, "home.tmpl", p)
}
