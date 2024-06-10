package main

import (
	"GinChat/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// db, err := gorm.Open(mysql.Open("root:123456@tcp(127.0.0.1:3301)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	dsn := "root:123456@tcp(localhost:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Name = "maut"
	db.Create(&user)
	// db.Create(&models.UserBasic{Name: "maut"})

	// Read
	// var product Product
	fmt.Printf("db.First(user, 1): %v\n", db.First(user, 1)) // find product with integer primary key
	// db.First(user, "code = ?", "D42")                        // find product with code D42

	// Update - update product's price to 200
	db.Model(user).Update("Password", 123456789)
	// Update - update multiple fields
	// db.Model(user).Updates(models.UserBasic{Price: 200, Code: "F42"}) // non-zero fields
	// db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	// db.Delete(user, 1)
}
