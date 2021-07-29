package router

import (
	"ginRouterDemo/helps"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

func Include(opts ...Option) { //引入路由
	options = append(options, opts...)
}

//初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()

	r.GET("/sn", sign) //加密路由

	for _, opt := range options {
		opt(r)
	}
	return r
}

//sign 加密
func sign(c *gin.Context) {
	nowTime := helps.GetTimeUnix()
	//int 转 string
	ts := strconv.FormatInt(nowTime, 10)

	//接受用户信息
	// user := c.PostForm("username")
	// pwd := c.PostForm("password")

	//声明一个map
	res := map[string]interface{}{}
	//set param
	params := url.Values{
		"name":  []string{"qwer"},
		"price": []string{"100"},
		"ts":    []string{ts},
	}
	res["sn"] = helps.CreateSign(params) //生成的签名
	res["ts"] = ts
	helps.ReturnJosn(c, "200", "", res)
}
