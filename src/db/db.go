package db

import (
	"../config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"
)

const (
	dbConnectionUrl string = "{login}:{password}@tcp({host}:{port})/{table}?charset=utf8&parseTime=True&loc=Local"
)

var db *gorm.DB

func InitDatabase() {
	connectionUrl := strings.Replace(dbConnectionUrl, "{login}", viper.GetString(config.DbLogin), -1)
	connectionUrl = strings.Replace(connectionUrl, "{password}", viper.GetString(config.DbPassword), -1)
	connectionUrl = strings.Replace(connectionUrl, "{host}", viper.GetString(config.DbHost), -1)
	connectionUrl = strings.Replace(connectionUrl, "{port}", viper.GetString(config.DbPort), -1)
	connectionUrl = strings.Replace(connectionUrl, "{table}", viper.GetString(config.DbName), -1)

	// Connect to db
	conn, err := gorm.Open("mysql", connectionUrl)

	if err != nil {
		panic(err)
	}

	// Test connection to database
	if err = conn.DB().Ping(); err != nil {
		panic(err)
	} else {
		log.Println("Connection to DB established")
	}

	// Set connection limitations
	conn.DB().SetMaxOpenConns(viper.GetInt(config.DbMaxOpenConn))
	conn.DB().SetMaxIdleConns(viper.GetInt(config.DbMaxIdleConn))
	conn.DB().SetConnMaxLifetime(time.Duration(viper.GetInt(config.DbMaxConnLifetime)) * time.Second)

	//
	db = conn
	//db.Debug().AutoMigrate(&entities.UserEntity{})
}

func GetDB() *gorm.DB {
	return db
}
