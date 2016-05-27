package util

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/inconshreveable/log15"
)

func BuildLocalTimeStr(utcTime time.Time) (string, error) {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return "N/A", err
	}
	return utcTime.In(loc).Format("15:04"), nil
}

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
