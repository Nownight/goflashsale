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
        return // 💡 必须加 return！报错后立刻踩刹车，停止执行下方的代码
    }
    result:=conf.DB.Create(&product)
    if result.Error !=nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"进货失败，仓库出现故障！"})
        return // 💡 必须加 return！
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
        return // 💡 必须加 return！
    }
    
    // 💡 修改这里：拆掉 "data" 包装盒，直接把纯数组扔给 React！
    // 这样你的 React 代码里的 products.map 才能完美运行
    c.JSON(http.StatusOK, products)
}