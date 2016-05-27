package webpage

import (
	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

//HomePage 首页
func HomePage(c *gin.Context) {
	log := c.MustGet("log").(log15.Logger)
	log.Info("展现首页处理开始")
	c.HTML(200, "home.tmpl", "")
}
