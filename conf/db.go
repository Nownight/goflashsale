// 文件夹：conf / 文件：db.go
package conf

import (
	"fmt"
	"os"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB 就像是通往仓库的“传送带”，全店共享这一条
var DB *gorm.DB

// InitDB 是装修工人接管线的动作
func InitDB() {
	// 这是仓库的地址和钥匙（账号、密码、数据库名）
	dsn := os.Getenv("DB_URL")
	if dsn == "" {
		// ⚠️ 绝密警告：传到 GitHub 的代码里，千万不能出现 TiDB 的真实密码！
		// 这里换回本地假地址占位，真正的云密码我们等会儿在 Render 后台悄悄配置。
		dsn = "root:20182019@tcp(127.0.0.1:3306)/goflashsale?charset=utf8mb4&parseTime=True&loc=Local"
	}
	// 尝试连接仓库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("💥 仓库连接失败，管线没接通: "+err.Error())
	}
	DB=db
	fmt.Println("✅ 仓库管线连接成功！")
}