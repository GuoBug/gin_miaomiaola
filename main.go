package main

import (
	"html/template"

	"gitlab.mentornow.com/leroad-service/core/common/handler"

	"gin_miaomiaola/util"
	"gin_miaomiaola/webpage"

	"github.com/gin-gonic/gin"
)

func main() {
	m := gin.Default()
	m.Use(handler.Log()) //log15日志

	funcs := template.FuncMap{
		"timeFmtDate": util.BuildLocalDateStr,
		"timeFmtTime": util.BuildLocalTimeStr,
	}

	templ := template.Must(template.New("projectViews").Funcs(funcs).ParseGlob("templates/*.tmpl"))

	m.SetHTMLTemplate(templ)

	m.GET("/", webpage.HomePage)
	m.Run(":3000")
}
