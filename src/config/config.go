package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

const (
	DbHost            = "db.host"
	DbPort            = "db.port"
	DbLogin           = "db.login"
	DbPassword        = "db.password"
	DbName            = "db.name"
	DbMaxOpenConn     = "db.conn.max_open_connections"
	DbMaxIdleConn     = "db.conn.max_idle_connections"
	DbMaxConnLifetime = "db.conn.max_connection_lifetime_seconds"

	AzureToken        = "azure.token"
	AzureOrganization = "azure.organization_name"
)

const (
	configPath      = "./"
	configName      = "config"
	configExtension = ".yml"
)

func LoadConfig() {
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)

	createDefaultConfig()

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	log.Println("Config loaded successfully")
}

// Creates default config if not exists
// https://gist.github.com/novalagung/13c5c8f4d30e0c4bff27
func createDefaultConfig() {
	var path = configPath + configName + configExtension
	// Detect if file exists
	var _, err = os.Stat(path)

	// Create default config file if not exists
	if os.IsNotExist(err) {
		log.Println("Config does not exist, creating")
		var file, err = os.Create(path)
		if isError(err) {
			return
		}
		defer file.Close()

		log.Println("Default config created in ", path)
		fillDefaultConfig()
		log.Println("Exiting as you should fill it")
		os.Exit(0)
	} else {
		log.Println("Config exists") // TODO Remove it
	}
}

func fillDefaultConfig() {
	// Db
	viper.SetDefault(DbHost, "localhost")
	viper.SetDefault(DbPort, 3306)
	viper.SetDefault(DbLogin, "login")
	viper.SetDefault(DbPassword, "password")
	viper.SetDefault(DbName, "databaseName")

	viper.SetDefault(DbMaxOpenConn, 5)
	viper.SetDefault(DbMaxIdleConn, 5)
	viper.SetDefault(DbMaxConnLifetime, 60*5)

	// Azure
	viper.SetDefault(AzureToken, "token")
	viper.SetDefault(AzureOrganization, "organizationName")

	err := viper.WriteConfig()
	if err != nil {
		panic(fmt.Errorf("%s", err))
	}

	log.Println("Current config params: ", viper.AllSettings())
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return err != nil
}
