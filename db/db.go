package db

import (
	"log"
	"os"
	"time"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	dbgorm *gorm.DB
	err    error
)

//ConfigureDB db configuration
func ConfigureDB() {
	dbgorm, err = gorm.Open("mysql", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
		log.Fatalln("Cannot connect to database")
	}

	log.Println("DB Connected")

	dbgorm.SingularTable(true)
	dbgorm.DB().SetMaxIdleConns(10)
	dbgorm.DB().SetMaxOpenConns(30)
	dbgorm.DB().SetConnMaxLifetime(time.Hour)
}

//GetDB return dbgorm
func GetDB() *gorm.DB {
	return dbgorm
}
