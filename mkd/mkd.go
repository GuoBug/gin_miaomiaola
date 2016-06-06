package mkd

import (
	"github.com/inconshreveable/log15"
	"github.com/russross/blackfriday"
)

//MarkdownToHTML Markdown 翻译
func MarkdownToHTML(input string, log log15.Logger) string {
	log.Debug("转换成的HTML")

	//input = "Hi here *is* a test"
	output := blackfriday.MarkdownBasic([]byte(input))
	return string(output)
}
