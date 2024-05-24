package ginaction

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	Server := gin.Default()

	Server.GET("/", func(context *gin.Context) {
		context.ClientIP()
	})

	AutoRegister(Server.Group("api"), TestDemo{})

	server := &http.Server{
		Addr:    ":8080",
		Handler: Server,
	}

	//启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	//等待一个INT或TERM信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	//创建超时上下文，Shutdown可以让未处理的连接在这个时间内关闭
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//停止HTTP服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	return
}

type TestDemo struct {
}

// ChooseMid 可以选择的服务中间件
func (a TestDemo) ChooseMid(t MidType) []gin.HandlerFunc {

	switch t {
	case 1:
		return []gin.HandlerFunc{Trace2()}
	default:
		return []gin.HandlerFunc{Trace()}
	}
}

// Group 接口组标识
func (a TestDemo) Group() string {
	return "test"
}

// Register 接口注册
func (a TestDemo) Register() []Action {
	return []Action{
		NewAction("GET", a.test1),
		NewAction("GET", a.testPath2, UseMidType(1), UseLastPath("ttt")),
		NewAction("POST", a.testPath3, UseMidType(1), UseMidSep(0)),
		NewAction("POST", a.testPath4, UseMidSep('_')),
	}
}

func (a TestDemo) test1(c *gin.Context) {
	time.Sleep(10 * time.Second)
	c.JSON(200, gin.H{
		"aaa": 1,
	})
}

func (a TestDemo) testPath2(c *gin.Context) {

}
func (a TestDemo) testPath3(c *gin.Context) {

}

func (a TestDemo) testPath4(c *gin.Context) {

}

// Trace 链路追踪
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Printf("12312dsaflk\r\n")
		c.Next()
	}
}

// Trace 链路追踪
func Trace2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println(222)
		c.Next()
	}
}
