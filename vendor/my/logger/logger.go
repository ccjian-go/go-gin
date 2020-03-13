package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func init(){
	fmt.Println("logger初始化中")
	// 创建记录日志的文件
	f, _ := os.Create("gin.log")

	//gin.DefaultWriter = io.MultiWriter(f)
	// 如果需要将日志同时写入文件和控制台，请使用以下代码
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	fmt.Println("logger初始化完毕")
}