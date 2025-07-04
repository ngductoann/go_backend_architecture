package benchmark

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID   uint
	Name string
}

func insertRecord(b *testing.B, db *gorm.DB) {
	user := User{Name: "ngductoann"}
	if err := db.Create(&user).Error; err != nil {
		b.Fatal(err)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	dsn := "toan:ductoan0712@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if you want
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// handle error if you want
			fmt.Println("Existing table found, dropping table", err)
		}
	}

	// Create table if not exists
	db.AutoMigrate(&User{})
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set args connection pool
	sqlDB.SetMaxIdleConns(1)
	defer sqlDB.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	dsn := "toan:ductoan0712@tcp(127.0.0.1:3306)/shopdevgo?charset=utf8mb4&parseTime=True&"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// SkipDefaultTransaction: false,
		// Logger:                 false,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Check if the table exists
	if db.Migrator().HasTable(&User{}) {
		// Drop the table if it exists
		if err := db.Migrator().DropTable(&User{}); err != nil {
			// Handle error if you want
			// fmt.Println("Existing table found, dropping table", err)
		}
	}

	// Create table if not exists
	db.AutoMigrate(&User{})
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// Set args connection pool
	sqlDb.SetMaxIdleConns(10)
	defer sqlDb.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
