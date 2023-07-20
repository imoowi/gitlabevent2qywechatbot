package middlewares

import (
	"gitlab2wechatbot/global"
	"gitlab2wechatbot/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GitlabTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-Gitlab-Token")
		if token == `` {
			response.Error(`pls set X-Gitlab-Token header `, http.StatusBadRequest, c)
			c.Abort()
			return
		}
		if token != global.Config.GetString(`webhook.token`) {
			response.Error(`token invalid `, http.StatusBadRequest, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
