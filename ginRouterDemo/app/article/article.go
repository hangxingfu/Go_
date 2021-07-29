package article

import (
	"ginRouterDemo/helps"

	"github.com/gin-gonic/gin"
)

func LoadArticle(r *gin.Engine) {
	// article
	article := r.Group("article", helps.VerifySign) //article路由组添加验证操作
	{
		article.GET("/index", articleIndex)
		article.GET("/content", articleContent)
	}
	// ...
}
