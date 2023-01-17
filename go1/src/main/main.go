package main

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "root:9264wkn.@tcp(localhost:3306)/example?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{SkipDefaultTransaction: true, PrepareStmt: true})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	product := &Product{}

	if err := db.First(product, 1).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO 业务处理查询不到数据的情况
			fmt.Println("没有找到")
		}
	}
	// find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}
