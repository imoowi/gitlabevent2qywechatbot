package swagger

import (
	"github.com/gin-gonic/gin"
	_ "gitlab2wechatbot/docs"
	"gitlab2wechatbot/router"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	router.RegisterRoute(Routers)
}

func Routers(e *gin.Engine) {
	// swagger
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
