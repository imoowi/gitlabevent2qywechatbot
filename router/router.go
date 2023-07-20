/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package router

import (
	"gitlab2wechatbot/middlewares"

	"github.com/gin-gonic/gin"
)

type Router func(*gin.Engine)

var router = []Router{}

func InitRouter() *gin.Engine {
	r := gin.Default()
	middlewares.InitMiddleware(r)
	for _, route := range router {
		route(r)

	}
	return r
}

// router  其余模块在init中调用RegisterRoute注册
func RegisterRoute(r ...Router) {
	router = append(router, r...)
}
