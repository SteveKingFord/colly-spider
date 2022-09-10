package initiallize

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//数据库配置
const (
	userName = "postgres"
	password = "fD52ki3qPbH21ORY6QyZYt6CCdUC"
	ip       = "bcore.top"
	port     = 5432
	dbName   = "spider"
)

func InitialGORM() *gorm.DB {
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, ip, port, dbName)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", ip, userName, password, dbName, port)

	fmt.Println("dsn:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("MySQL启动异常:", err.Error())
		return nil
	}

	return db
}
