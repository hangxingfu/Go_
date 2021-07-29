# Go 框架 Gin 的路由

## 主要内容 路由注册，分组，加token验证

整体目录结构如下，app目录下是项目的所有路由文件，而router.go文件是用来引入app下的路由文件和注册路由的。app下的每一个子目录就是一组路由。

```
.
├── app
│   ├── article
│   │   ├── article.go
│   │   └── handle.go
│   └── blog
│       ├── blog.go
│       └── handle.go
├── config
│   └── config.go
├── go.mod
├── go.sum
├── helps
│   └── func.go
├── main.go
└── router
    └── router.go
```

路由验证方面使用MD5加密参数和设置时间是否过期的方式来验证的。代码大致如下

```go
//生成签名
func CreateSign(param url.Values) string {
	var key []string //声明一个数组
	var str string

	//将传来的 键 参数存入key
	for v := range param {
		if v != "sn" {
			key = append(key, v)
		}
	}
	sort.Strings(key) // 排序
	// 循环拼接加密字串
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str = fmt.Sprintf("%v=%v", key[i], param.Get(key[i]))
		} else {
			str = str + fmt.Sprintf("&%v=%v", key[i], param.Get(key[i]))
		}
	}
	//md5加密
	return MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
}

//验证签名
func VerifySign(c *gin.Context) {
	var method = c.Request.Method //请求方式
	var formData url.Values       //表单数据
	var ts int64                  //验证时间
	var sn string                 //加密字符串

	if method == "GET" {
		formData = c.Request.URL.Query()
		ts, _ = strconv.ParseInt(c.Query("ts"), 10, 64) //将string转int64，用作对比时间是否过期
		sn = c.Query("sn")
	} else if method == "POST" {
		c.Request.ParseForm()
		formData = c.Request.PostForm
		ts, _ = strconv.ParseInt(c.PostForm("ts"), 10, 64)
		sn = c.PostForm("sn")
	} else {
		ReturnJosn(c, "500", "Error: Illegal Request", "")
		return
	}

	//验证过期时间
	if ts > GetTimeUnix() || GetTimeUnix()-ts > config.API_EXPIRY {
		ReturnJosn(c, "500", "Error: timestamp expiry", "")
		return
	}
	//验证签名
	if sn == "" || sn != CreateSign(formData) {
		ReturnJosn(c, "500", " Error: sign", "")
		return
	}

}
```

>此处代码，参考 `https://github.com/xinliangnote/Go` && `https://www.liwenzhou.com/posts/Go/gin_routes_registry/`

---

使用mod工具添加依赖，clone项目后请先下载依赖再运行
