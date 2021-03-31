package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/logger"
	"web_app/routes"
	"web_app/settings"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	var confPath string

	flag.StringVar(&confPath, "path", "", "配置文件路径")
	flag.Parse()
	if confPath == "" {
		fmt.Println("config path miss")
		return
	}
	//初始化配置文件
	settings.Init(confPath)
	//初始化日志
	logger.Init(settings.Conf.LogConf)
	defer zap.L().Sync()
	//初始化mysql
	mysql.Init(settings.Conf.DbConf)
	defer mysql.Close()
	//初始化redis
	redis.Init(settings.Conf.RedisConf)
	defer redis.Close()
	//初始化路由
	r := routes.Setup()
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	zap.L().Info("shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Fatal("server shutdown", zap.Error(err))
	}
	zap.L().Info("server exiting")
}
