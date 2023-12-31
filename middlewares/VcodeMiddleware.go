package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gitlab2wechatbot/components"
	"gitlab2wechatbot/utils/response"
)

type POSTVcode struct {
	Id    string `json:"captcha_id" form:"captcha_id" binding:"required"`
	Vcode string `json:"captcha_code" form:"captcha_code" binding:"required"`
}

func VcodeMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		var postVcode *POSTVcode
		err := c.ShouldBindBodyWith(&postVcode, binding.JSON)
		if err != nil {
			response.Error(err.Error(), http.StatusBadRequest, c)
			c.Abort()
			return
		}
		if !components.VerifyCaptcha(postVcode.Id, postVcode.Vcode) {
			response.Error(`验证码验证失败`, http.StatusBadRequest, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
