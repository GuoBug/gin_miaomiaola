package main

import (
	"html/template"
	"net/http"

	"gin_miaomiaola/util"
	"gin_miaomiaola/webpage"

	"github.com/gin-gonic/gin"
)

func main() {
	m := gin.Default()
	m.Use(util.Log()) //log15日志
	m.Use(util.DB())

	funcs := template.FuncMap{
		"timeFmtDate": util.BuildLocalDateStr,
		"timeFmtTime": util.BuildLocalTimeStr,
	}

	templ := template.Must(template.New("projectViews").Funcs(funcs).ParseGlob("templates/*.tmpl"))

	m.StaticFS("/static", http.Dir("./static"))
	m.SetHTMLTemplate(templ)

	m.GET("/", webpage.HomePage)
	m.GET("/post/:urlString", webpage.PostPage)

	m.Run(":4000")
}
