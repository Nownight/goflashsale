package routes

import (
    "goflashsale/controllers"
    "goflashsale/middleware"
    "github.com/gin-gonic/gin"
)

// NewRouter 负责规划整家店的路线图
func SetupRoutes(r *gin.Engine) {

    // 【公共区域】：大门口，谁都能来
    r.POST("/register", controllers.Register) 
    r.POST("/login", controllers.Login)       
    
    // 💡 关键修改：把看菜单的接口挪到公共区域！
    // 这样前端 fetch("/products") 就不会 404，也不需要 token 啦！
    r.GET("/products", controllers.GetProducts) 

    // 【VIP 区域】：必须过保安那一关
    auth := r.Group("/api")           
    auth.Use(middleware.JWTAuth())    
    {
        auth.GET("/me", func(c *gin.Context) {
            id, _ := c.Get("user_id") 
            c.JSON(200, gin.H{"message": "进内场了！", "your_id": id})
        })
        
        // 进货（商家操作，需要内场权限）
        auth.POST("/product", controllers.CreateProduct)
        
        // 顾客下单、看自己的订单记录（必须登录身份）
        auth.POST("/order", controllers.CreateOrder)
        auth.GET("/orders", controllers.GetMyOrders)
    }
}