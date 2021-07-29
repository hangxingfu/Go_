package article

import (
	"ginRouterDemo/helps"

	"github.com/gin-gonic/gin"
)

func articleIndex(c *gin.Context) {
	helps.ReturnJosn(c, "200", "this is a article index.", "")
}

func articleContent(c *gin.Context) {
	helps.ReturnJosn(c, "200", "this is a article content.", "")
}

// fun ...
