package models

import "gorm.io/gorm"

// Product 就是咱们店里的“商品登记表”
type Product struct {
	gorm.Model // 老规矩：自带商品编号(ID)、上架时间等隐藏字段

	Name        string  // 商品名称（比如：招牌珍珠奶茶）
	Description string  // 商品描述（比如：好喝到爆的绝版奶茶）
	Price       float64 // 价格（用 float64 是因为钱有小数，比如 9.9 元）
	Stock       int     // 库存数量（重点！秒杀全靠它，卖完就不能再卖了）
}