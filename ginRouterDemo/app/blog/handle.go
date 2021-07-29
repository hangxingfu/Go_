package blog

import (
	"ginRouterDemo/helps"

	"github.com/gin-gonic/gin"
)

func blogIndex(c *gin.Context) {
	helps.ReturnJosn(c, "200", "this is a blog index.", "")
}

func blogDetial(c *gin.Context) {
	helps.ReturnJosn(c, "200", "this is a blog detail.", "")
}
