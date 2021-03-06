package database
import (
	"assignment-2/models"
	"fmt"
	"log"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	host     = "localhost"
	user     = "root"
	password = ""
	dbPort   = "443"
	dbName   = "mygram"
	DB       *gorm.DB
	err      error
)

func MulaiDB() {
	dsn := "root@tcp(127.0.0.1:3306)/orders_by?charset=utf8mb4&parseTime=True&loc=Local"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal menyambung ke database :", err)
	}

	fmt.Println("Koneksi Sukses")
	database.Debug().AutoMigrate(models.Order{}, models.Item{})
	DB = database
}

func AmbilDB() *gorm.DB {
	return DB
}