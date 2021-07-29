package blog

import (
	"github.com/gin-gonic/gin"
)

func LoadBlog(r *gin.Engine) {
	blog := r.Group("blog")
	{
		blog.GET("/index", blogIndex)
		blog.GET("/detail", blogDetial)
	}
}
