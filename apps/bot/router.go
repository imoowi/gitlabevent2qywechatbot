/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package bot

import (
	"gitlab2wechatbot/apps/bot/handlers"
	"gitlab2wechatbot/middlewares"
	"gitlab2wechatbot/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	api := e.Group("/api")
	bots := api.Group("/bots")
	bots.Use(middlewares.GitlabTokenMiddleware())
	{
		bots.POST("", handlers.BotAdd) //新增
	}

	//!import:do-not-delete-this-line,不要删除此行，主要用于代码生成器
}
