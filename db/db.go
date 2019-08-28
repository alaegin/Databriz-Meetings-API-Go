package db

import (
	"Databriz-Meetings-API-Go/configs"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	dbConnectionUrl string = "{login}:{password}@tcp({host}:{port})/{table}?charset=utf8&parseTime=True&loc=Local"
)

var db *gorm.DB

func InitDatabase() {
	connectionUrl := strings.Replace(dbConnectionUrl, "{login}", viper.GetString(configs.DbLogin), -1)
	connectionUrl = strings.Replace(connectionUrl, "{password}", viper.GetString(configs.DbPassword), -1)
	connectionUrl = strings.Replace(connectionUrl, "{host}", viper.GetString(configs.DbHost), -1)
	connectionUrl = strings.Replace(connectionUrl, "{port}", viper.GetString(configs.DbPort), -1)
	connectionUrl = strings.Replace(connectionUrl, "{table}", viper.GetString(configs.DbName), -1)

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
	conn.DB().SetMaxOpenConns(viper.GetInt(configs.DbMaxOpenConn))
	conn.DB().SetMaxIdleConns(viper.GetInt(configs.DbMaxIdleConn))
	conn.DB().SetConnMaxLifetime(time.Duration(viper.GetInt(configs.DbMaxConnLifetime)) * time.Second)

	//
	db = conn
	//db.Debug().AutoMigrate(&entities.UserEntity{})
}

func GetDB() *gorm.DB {
	return db
}
