package main

import (
	"blog_tpl/config"
	"blog_tpl/daos"
	"blog_tpl/logger"
	"blog_tpl/routers"
	"blog_tpl/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
项目结构
config ---> 配置文件
controllers -> 控制器层 处理请求
logger ----> 日志处理
daos ---> 数据层
	mysql --> 数据获取
	redis --->
	es
logics ---> 逻辑层
middlewares  ---> 中间件层
models ---> 模型层
routers ---> 路由配置
utils ---> 工具类
*/
func main() {

	//初始化配置
	config.Init()

	//初始化全局ID生成器
	utils.Init(uint16(viper.GetInt("app.machine_id")))

	//初始化日志库
	logger.InitLogger()
	//刷新缓冲区
	defer zap.L().Sync()
	//zap.L().Info("init mysql success")
	//zap.L().Error("init redis success ", zap.Error(errors.New("fail fail fail")))

	//初始化dao层数据
	daos.InitDaos()

	//加载路由信息
	appRouter := routers.SetupRouter()

	//appRouter.Run()

	//启动服务 并添加优雅关闭
	elegantCloseApp(appRouter)

}

//优雅关闭
func elegantCloseApp(appRouter *gin.Engine) {

	genid, _ := utils.GetID()
	fmt.Println("-----------", genid)

	//优雅关闭
	srv := &http.Server{
		Addr:    ":8080",
		Handler: appRouter,
	}
	go func() {
		// 开启⼀一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号来优雅地关闭服务器器，为关闭服务器器操作设置⼀一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建⼀一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常⽤用的Ctrl+C就是触发系统SIGINT信号
	//kill -9 发送 syscall.SIGKILL 信号，但是不不能被捕获，所以不不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	//此处不不会阻塞
	//监听两个信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	// 阻塞在此，当接收到上述两种信号时才会往下执⾏行行
	<-quit
	log.Println("Shutdown Server ...")
	// 创建⼀一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务(将未处理理完的请求处理理完再关闭服务)，超过5秒就超时退出
	//相当于告诉程序给你5s的时间，把未完成的请求处理一下
	log.Println("****************** ...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}
	log.Println("Server exiting")
}
