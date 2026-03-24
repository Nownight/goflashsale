// 文件夹：middleware / 文件：jwt.go
package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("my_ultra_secret_key_123") // 保安手里也有一份私章的图纸，用来验真伪

// JWTAuth 就是那个铁面无私的保安
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 保安大喊：“出示手环！” (获取 Token)
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "站住！没带手环不准进"})
			c.Abort() // 拦住，不准往后走
			return
		}

		// 2. 保安仔细查验手环的真伪
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		// 3. 如果手环是假的，或者过期了
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "手环无效，滚出去！"})
			c.Abort()
			return
		}

		// 4. 查验通过！保安把手环里的 VIP 编号抄下来，贴在顾客背后，方便里面的人认出他
		claims := token.Claims.(jwt.MapClaims)
		c.Set("user_id", claims["user_id"])

		// 5. 放行！
		c.Next()
	}
}