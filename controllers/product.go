package controllers

import(
	"goflashsale/conf"
	"goflashsale/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context){
	var product models.Product
	if err:=c.ShouldBindJSON(&product); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"货单填写格式不对，拒收！"})

	}
	result:=conf.DB.Create(&product)
	if result.Error !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"进货失败，仓库出现故障！"})
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"新奶茶进口成功！",
		"product_id":product.ID,
	})
}

func GetProducts(c *gin.Context){
	var products []models.Product
	result:=conf.DB.Find(&products)
	if result.Error!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"获取商品列表失败！"})
	}
	c.JSON(http.StatusOK,gin.H{
		"data":products,
	})
}