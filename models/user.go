// 文件夹：models / 文件：user.go
package models

import "gorm.io/gorm"

// User 就是那张“会员登记表”的模板
type User struct {
	gorm.Model // 这行自带 ID（会员编号）、注册时间等隐藏字段

	// 要求顾客填写的三个格子
	UserName string `gorm:"unique"` // 名字（规定：全店不能有重名）
	Password string                 // 密码（存在仓库里的是加密后的乱码）
	Phone    string                 // 手机号
}