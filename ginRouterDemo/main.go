package main

import (
	"fmt"
	"ginRouterDemo/app/article"
	"ginRouterDemo/app/blog"
	"ginRouterDemo/router"
)

func main() {
	//引入路由
	router.Include(blog.LoadBlog, article.LoadArticle)
	//注册路由
	r := router.InitRouter()
	if err := r.Run(":8000"); err != nil { //启动服务
		fmt.Printf("gin services start failed , err:" + err.Error())
	}
}
