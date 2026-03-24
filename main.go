// 文件：main.go (在项目最外层)
package main

import (
	"goflashsale/conf"
	"goflashsale/models"
	"goflashsale/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
		if method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	// 1. 店长：接通仓库管线
	conf.InitDB()

	// 2. 店长：打造实体存放架
	conf.DB.AutoMigrate(&models.User{})
	conf.DB.AutoMigrate(&models.Product{})
	conf.DB.AutoMigrate(&models.Order{})

	// 3. 店长：建立大门
	r := gin.Default()

	// 🎯 4. 店长：把跨域保安安排在最前面！（极其重要）
	r.Use(Cors())

	// 🎯 5. 店长：把大门交给设计员去挂指示牌
	routes.SetupRoutes(r)

	// 🎯 6. 获取云平台分配的端口（如果是本地运行，默认用 8080）
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// 7. 正式开业！
	r.Run(":" + port)
}