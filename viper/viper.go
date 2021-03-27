package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

const MysqlIpKey = "mysql.ip"
const RedisIpKey = "redis.ip"

func testJSON(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	//Can be called multiple times to define multiple search paths
	viper.AddConfigPath(path + "/viper/json/")
	err1 := viper.ReadInConfig()
	if err1 != nil {
		log.Fatalf("load json/mysql.json error")
	}
	fmt.Println("mysql.ip : ", viper.Get(MysqlIpKey))
	fmt.Println("redis.ip : ", viper.Get(RedisIpKey))
}

func testTOML(path string) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(path + "/viper/toml/")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("load toml/setting.toml error")
	}
	fmt.Println("toml mysql.ip : ", viper.Get(MysqlIpKey))
	fmt.Println("toml redis.ip : ", viper.Get(RedisIpKey))
}

func init() {
	log.Println("viper init ...")
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd fault")
	}
	testJSON(path)
	//testTOML(path)
	time.Sleep(time.Second * 20)
	testJSON(path)
}
