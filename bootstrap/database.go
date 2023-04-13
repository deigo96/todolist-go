package bootstrap

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabaseConnection(env *Env) *gorm.DB {
	// db, err := sql.Open("mysql", "project:brengsek96@tcp(127.0.0.1:3306)/test")
	// dsn := "project:brengsek96@tcp(127.0.0.1:3306)/todolist?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := env.DbUser + ":" + env.DbPass + "@tcp(" + env.DbHost + ":" + env.DbPort + ")/" + env.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}

func CloseConnection(db *gorm.DB) {
	if db == nil {
		return
	}

	if db != nil {
		db, _ := db.DB()
		db.Close()
	}
}
