package controllers

import (
	"goflashsale/conf"
	"goflashsale/models"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func CreateOrder(c *gin.Context){
	var input struct{
		ProductID uint `json:"product_id"`
	}
	if err:=c.ShouldBindJSON(&input);err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"请选择商品",
		})
		return
	}

	userID,_:=c.Get("user_id")
	tx:=conf.DB.Begin()
	var product models.Product
	if err:=tx.Clauses(clause.Locking{Strength:"UPDATE"}).Where("id=?",input.ProductID).First(&product).Error; err!=nil{
		tx.Rollback()
		c.JSON(http.StatusNotFound,gin.H{"error":"商品不存在！"})
		return
	}
	if product.Stock<=0{
		tx.Rollback()
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"已售罄，下次早来！",
		})
		return
	}

	tx.Model(&product).Update("stock",product.Stock-1)
	order:=models.Order{
		UserID:uint(userID.(float64)),
		ProductID:input.ProductID,
	}
	tx.Create(&order)
	tx.Commit()

	c.JSON(http.StatusOK,gin.H{
		"message":"抢购成功！",
		"order_id":order.ID,
	})

}

func GetMyOrders(c *gin.Context){
	userID,_:=c.Get("user_id")
	var orders []models.Order
	result:=conf.DB.Preload("Product").Where("user_id",uint(userID.(float64))).Find(&orders)
	if result.Error!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"error":"获取订单失败！",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"my_orders":orders,
	})

}