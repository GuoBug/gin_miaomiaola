package util

import (
	"log"
	"os"
	"time"

	"gopkg.in/mgo.v2"

	"mentornow.com/leroad/general-api/conf"
	"mentornow.com/leroad/general-api/util"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

//DB 创建mongo session
func DB() gin.HandlerFunc {
	session, err := util.CreateDBSession()
	if err != nil {
		log.Panic("创建session失败", err)
	}
	return func(c *gin.Context) {
		s := session.Clone()
		c.Set("db", s.DB("miao_blog"))
		defer s.Close()
		c.Next()
	}
}

//CreateDBSession 创建db session
func CreateDBSession() (*mgo.Session, error) {
	log := log15.New()

	session, err := mgo.Dial(conf.Cfg.MongoURL)
	if err != nil {
		log.Error("创建session失败", err)
		return nil, err
	}
	admindb := session.DB(conf.Cfg.MongoAuthDB)
	err = admindb.Login(conf.Cfg.MongoUser, conf.Cfg.MongoPsw)
	if err != nil {
		log.Error("登陆数据库失败", err)
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}

//BuildLocalTimeStr time
func BuildLocalTimeStr(utcTime time.Time) (string, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "N/A", err
	}
	return utcTime.In(loc).Format("15:04"), nil
}

//BuildLocalDateStr date
func BuildLocalDateStr(utcTime time.Time) (string, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "N/A", err
	}

	return utcTime.In(loc).Format("01月02日"), nil
}

//Log 日志
func Log() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("X-Token")

		log := log15.New("Method", context.Request.Method, "Api", context.Request.URL.Path, "IP", context.ClientIP(), "Token", token)
		var handler log15.Handler
		if gin.Mode() == gin.DebugMode {
			handler = log15.StreamHandler(os.Stdout, log15.TerminalFormat())
			//handler = log15.StreamHandler(os.Stdout, log15.JsonFormatEx(true, true))
		} else {
			handler = log15.MultiHandler(
				log15.StreamHandler(os.Stdout, log15.TerminalFormat()),
				log15.LvlFilterHandler(log15.LvlError, log15.Must.FileHandler("/var/log/web/error.log", log15.JsonFormat())),
				log15.LvlFilterHandler(log15.LvlInfo, log15.Must.FileHandler("/var/log/web/info.log", log15.JsonFormat())),
			)
		}

		h := log15.CallerFileHandler(handler)
		log.SetHandler(h)

		context.Set("log", log)
		context.Next()

	}
}
