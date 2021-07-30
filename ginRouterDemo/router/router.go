package router

import (
	"ginRouterDemo/helps"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/fvbock/endless"
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

	//平滑重启
	r.GET("/", func(c *gin.Context) {
		// time.Sleep(time.Second * 15)
		c.String(http.StatusOK, "佐贺偶像是传奇\n夜露死苦！！！！！！SAGA~   ~~\ntest graceful_restart\ntest gogogogoggo")
	})

	r.GET("/sn", sign) //加密路由

	for _, opt := range options {
		opt(r)
	}

	/*
		默认endless服务器会监听下列信号：
		syscall.SIGHUP，syscall.SIGUSR1，syscall.SIGUSR2，syscall.SIGINT，syscall.SIGTERM和syscall.SIGTSTP
		接收到 SIGHUP 信号将触发`fork/restart` 实现优雅重启（kill -1 pid会发送SIGHUP信号）
		接收到 syscall.SIGINT或syscall.SIGTERM 信号将触发优雅关机
		接收到 SIGUSR2 信号将触发HammerTime
		SIGUSR1 和 SIGTSTP 被用来触发一些用户自定义的hook函数
	*/
	if err := endless.ListenAndServe(":8000", r); err != nil {
		log.Fatalf("listen error:%s\n", err)
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
