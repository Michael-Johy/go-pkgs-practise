package models

import (
	"fmt"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/setting"
	"gorm.io/gorm"
	"log"
)
import "gorm.io/driver/mysql"

var DB *gorm.DB

func init() {
	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/gin_test?charset=utf8&parseTime=True&loc=Local", setting.MysqlUser, setting.MysqlPassword,
		setting.MysqlIP, setting.MysqlPort)

	var err error

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dns,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("open mysql fault")
	}

}
