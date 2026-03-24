// 文件夹：routes / 文件：router.go
package routes

import (
	"goflashsale/controllers"
	"goflashsale/middleware"
	"github.com/gin-gonic/gin"
)

// NewRouter 负责规划整家店的路线图
func SetupRoutes(r *gin.Engine) {

	// 【公共区域】：谁都能来
	r.POST("/register", controllers.Register) // 牌子：注册 -> 指向注册员
	r.POST("/login", controllers.Login)       // 牌子：登录 -> 指向发证员

	// 【VIP 区域】：必须过保安那一关
	auth := r.Group("/api")           // 划定一个内场区域
	auth.Use(middleware.JWTAuth())    // 把保安安排在内场门口
	{
		// 牌子：查看个人信息 -> 只有带了手环进了内场的人才能走到这里
		auth.GET("/me", func(c *gin.Context) {
			id, _ := c.Get("user_id") // 办事员看一眼保安贴的便利贴
			c.JSON(200, gin.H{"message": "进内场了！", "your_id": id})
		})
		auth.POST("/product",controllers.CreateProduct)
		auth.GET("/products",controllers.GetProducts)
		auth.POST("/order",controllers.CreateOrder)
		auth.GET("/orders",controllers.GetMyOrders)
	}
}