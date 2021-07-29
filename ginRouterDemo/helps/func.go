package helps

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"ginRouterDemo/config"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ReturnJosn(c *gin.Context, code, str string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": str,
		"data":    data,
	})
	c.Abort()
}

//获取当前时间戳
func GetTimeUnix() int64 {
	return time.Now().Unix()
}

//md5加密
func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

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
