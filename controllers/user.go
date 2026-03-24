// 文件夹：controllers / 文件：user.go
package controllers

import (
	"goflashsale/conf"
	"goflashsale/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// 制卡机的私章（全店最高机密）
var jwtSecret = []byte("my_ultra_secret_key_123")

// -----------------------------------------
// 办事员 1 号：负责给新顾客办理注册
// -----------------------------------------
func Register(c *gin.Context) {
	var user models.User // 拿出一张空表格

	// 1. 检查顾客递交的资料格式对不对
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "资料填写格式不对！"})
		return // 格式不对，直接赶走
	}

	// 2. 为了安全，把明文密码放进碎纸机，变成乱码
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword) // 把乱码填回表格

	// 3. 把填好的表格扔进传送带，存入仓库
	result := conf.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "注册失败，名字可能被别人抢注了"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功！", "id": user.ID})
}

// -----------------------------------------
// 办事员 2 号：负责老顾客登录，并颁发“电子手环”
// -----------------------------------------
func Login(c *gin.Context) {
	var inputUser models.User // 顾客递交的账号密码
	var dbUser models.User    // 从仓库里找出来的真实档案

	// 1. 收取顾客递交的账号密码
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "输入格式不对！"})
		return
	}

	// 2. 去仓库里查有没有这个人
	result := conf.DB.Where("user_name = ?", inputUser.UserName).First(&dbUser)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "查无此人！"})
		return
	}

	// 3. 把顾客输入的密码，和仓库里的乱码进行比对
	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(inputUser.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "密码错误！"})
		return
	}

	// 4. 比对成功！启动制卡机，给顾客发一个电子手环 (Token)
	token, _ := GenerateToken(dbUser.ID)
	
	c.JSON(http.StatusOK, gin.H{
		"message": "登录成功！请戴好手环",
		"token":   token,
	})
}

// -----------------------------------------
// 内部工具：制卡机（制作手环）
// -----------------------------------------
func GenerateToken(userID uint) (string, error) {
	// 把顾客的 ID 刻在手环上，有效期 24 小时
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret) // 盖上私章，出卡！
}