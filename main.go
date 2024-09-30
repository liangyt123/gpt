package main

import (
	"mygame/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// CORS 设置
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 允许所有来源
		AllowMethods:     []string{"GET", "POST"},                             // 允许的请求方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12小时的有效期
	}))
	// 初始化玩家信息
	// Serve static files
	router.Static("/static", "./static")

	// Routes
	router.POST("/api/player", controllers.GetPlayerInfo)
	router.POST("/api/choose", controllers.MakeChoice)

	// Serve HTML
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.Run(":80") // 监听8080端口
}
