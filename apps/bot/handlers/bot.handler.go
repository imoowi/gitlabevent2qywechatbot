/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package handlers

import (
	"gitlab2wechatbot/apps/bot/services"
	"gitlab2wechatbot/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary	新增(add)
// @Tags		bot
// @Accept		application/json
// @Produce	application/json
// @Success	200
// @Failure 400                        "请求错误"
// @Failure 401                        "token验证失败"
// @Failure 500                         "内部错误"
// @Router		/api/bots [post]
func BotAdd(c *gin.Context) {
	// var bot *models.GitLabEvent
	var bot map[string]any
	err := c.ShouldBindJSON(&bot)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	services.Insert2Chan(bot)
	// for i := 0; i < 60; i++ {
	// 	go services.Insert2Chan(bot)
	// }
	// fmt.Println(`end`)
	response.OK(`ok`, c)
}
