package setting

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var (
	MysqlIP       string
	MysqlPort     int
	MysqlUser     string
	MysqlPassword string

	RedisIp       string
	RedisPort     int
	RedisPassword string

	HTTPPORT     int
	ReadTimeout  int
	WriteTimeout int

	PageSize = 5
)

func init() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd fault")
	}
	viper.SetConfigName("setting")
	viper.SetConfigType("toml")
	//Can be called multiple times to define multiple search paths
	viper.AddConfigPath(path + "/gin/conf/")
	err1 := viper.ReadInConfig()
	if err1 != nil {
		log.Fatalf("load conf/setting.toml error")
	}
	MysqlIP = viper.GetString("mysql.ip")
	MysqlPort = viper.GetInt("mysql.port")
	MysqlUser = viper.GetString("mysql.user")
	MysqlPassword = viper.GetString("mysql.password")

	RedisIp = viper.GetString("redis.ip")
	RedisPort = viper.GetInt("redis.port")
	RedisPassword = viper.GetString("redis.password")

	HTTPPORT = viper.GetInt("http.port")
	ReadTimeout = viper.GetInt("http.readTimeout")
	WriteTimeout = viper.GetInt("http.writeTimeout")

}
