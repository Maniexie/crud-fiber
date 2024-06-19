package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/crud_go?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := MYSQL

	DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if err != nil {
		panic("Database tidak Konek")
	}
	fmt.Println("Koneksi Berhasil")

	// DSN := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
}
