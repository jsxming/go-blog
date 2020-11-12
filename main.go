package main

import (
	"blog/global"
	"blog/internal/api"
	"blog/internal/model"
	"blog/pkg"
	"blog/pkg/util"
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // mysql
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/**
初始化配置文件的加载
*/
func setupSetting() error {
	setting, err := pkg.NewSetting()
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeOut *= time.Second
	global.ServerSetting.WriteTimeOut *= time.Second

	return nil
}

/**
初始化数据库
*/
func setupDBEngine() error {
	var err error
	err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	path := global.AppSetting.LogSavePath
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil || os.IsNotExist(err) {
		file, err = os.Create(path)
		return err
	}
	Log := logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: util.TimeLayout,
	})
	Log.SetOutput(file)
	global.Logger = Log
	return nil
}

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}
}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := api.NewRouter()

	s := &http.Server{
		Addr:           ":9620",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	//gin.New()

	//err := s.ListenAndServe()
	//if err != nil {
	//	fmt.Println("Server start error")
	//}

	// ctrl+c时 接收完请求后再 退出
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
