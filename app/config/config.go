package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var JWT_SECRRET = ""

type AppConfig struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string
	jwtKey     string
	KEY_API 			string
	KEY_API_SECRET 		string
	CLOUD_NAME			string
	IP_Public_Redis string
	Pass_Redis string
}

func InitConfig() *AppConfig {
	return ReadENV()
}

func ReadENV() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("JWT_KEY"); found {
		app.jwtKey = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBUSER"); found {
		app.DBUsername = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPASS"); found {
		app.DBPassword = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBHOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("DBPORT"); found {
		conv, _ := strconv.Atoi(val)
		app.DBPort = conv
		isRead = false
	}
	if val, found := os.LookupEnv("DBNAME"); found {
		app.DBName = val
		isRead = false
	}
	if val, found := os.LookupEnv("KEY_API"); found {
		app.KEY_API = val
		isRead = false
	}
	if val, found := os.LookupEnv("KEY_API_SECRET"); found {
		app.KEY_API_SECRET = val
		isRead = false
	}	
	if val, found := os.LookupEnv("CLOUD_NAME"); found {
		app.CLOUD_NAME = val
		isRead = false
	}
	if val, found := os.LookupEnv("IP_REDIS"); found {
		app.IP_Public_Redis = val
		isRead = false
	}
	if val, found := os.LookupEnv("PASS_REDIS"); found {
		app.Pass_Redis = val
		isRead = false
	}
	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		// viper.SetConfigName("server")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config: ", err.Error())
			return nil
		}
		app.jwtKey = viper.Get("JWT_KEY").(string)
		app.DBUsername = viper.Get("DBUSER").(string)
		app.DBPassword = viper.Get("DBPASS").(string)
		app.DBHost = viper.Get("DBHOST").(string)
		app.DBPort, _ = strconv.Atoi(viper.Get("DBPORT").(string))
		app.DBName = viper.Get("DBNAME").(string)
		app.KEY_API = viper.Get("KEY_API").(string)
		app.KEY_API_SECRET = viper.Get("KEY_API_SECRET").(string)
		app.CLOUD_NAME = viper.Get("CLOUD_NAME").(string)
		app.IP_Public_Redis = viper.Get("IP_REDIS").(string)
		app.Pass_Redis = viper.Get("PASS_REDIS").(string)
	}
	JWT_SECRRET = app.jwtKey
	return &app
}
