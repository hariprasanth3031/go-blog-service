package config

import (
	"fmt"
	mysql "go.elastic.co/apm/module/apmgormv2/v2/driver/mysql"
	"gorm.io/gorm"
	//"gorm.io/gorm/schema"
	// "time"
)

// PostgresDB : Variable to hold the postgresql db connection
var MariaDB *gorm.DB

// InitializeDB : Initializes the Database connections
func InitializeDB() {

	fmt.Println("env", Env.Dbconn)
	dsn := fmt.Sprintf("%s?parseTime=True", Env.Dbconn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	// sqlDB, _ := db.DB()
	// sqlDB.SetMaxIdleConns(25)
	// sqlDB.SetConnMaxLifetime(time.Minute * 15)

	MariaDB = db
}
