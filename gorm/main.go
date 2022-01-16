package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("conclusion.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	//db.First(&product, 1)                 // 根据整形主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	//fmt.Println("res", product)

	//// Update - 将 product 的 price 更新为 200
	//db.Model(&product).Update("Price", 200)
	//// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	//
	//// Delete - 删除 product
	//db.Delete(&product, 1)

	//将结构体传入where
	//type Para struct {
	//	Code  string
	//	Price uint
	//}

	//in := &Para{Code: "D42"}
	//out := &Para{}
	//db.Where("Code = ?", "D42").Find(&product)
	//fmt.Println("struct测试 ", product)

	pro := &Product{Code: "", Price: 100}
	// Struct
	db.Where(pro).First(&product)
	fmt.Println("struct测试1 ", product)
	//// SELECT * FROM users WHERE name = "jinzhu" AND age = 20 LIMIT 1;

	1234
}
