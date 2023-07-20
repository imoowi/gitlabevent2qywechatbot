/*
Copyright © 2023 yuanjun<simpleyuan@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"gitlab2wechatbot/apps/bot/services"
	"gitlab2wechatbot/global"
	"gitlab2wechatbot/router"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:          "server",
	Short:        "start api server",
	Example:      "gitlab2wechatbot server -c config/settings.yml",
	SilenceUsage: true,
	PreRun: func(cmd *cobra.Command, args []string) {
		setUp()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return run()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func setUp() {
	initServerConfig()
}
func initServerConfig() {
	global.Config.SetDefault("server.host", "0.0.0.0")
	global.Config.SetDefault("server.port", 8000)
}

func run() error {
	startServer()
	return nil
}

func startServer() {
	mode := global.Config.GetString("application.mode")
	switch mode {
	case "dev":
		gin.SetMode(gin.DebugMode)
	case "prod":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	}

	global.Bootstrap()
	r := router.InitRouter()

	port := global.Config.GetUint16("server.port")
	host := global.Config.GetString("server.host")
	s := &http.Server{
		Addr:           fmt.Sprintf(`%s:%d`, host, port),
		Handler:        r,
		ReadTimeout:    time.Duration(global.Config.GetInt("server.readtimeout")) * time.Second,
		WriteTimeout:   time.Duration(global.Config.GetInt("server.writertimeout")) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	go func() {
		services.LoopReadMsgChan()
	}()
	fmt.Println(`server port: `, port)
	fmt.Printf(`API document address http://localhost:%d/swagger/index.html`, port)
	fmt.Println()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(`Server Shutdown:`, err)
	}
	log.Println(`Server exiting`)
}