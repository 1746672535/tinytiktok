package srv

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"tinytiktok/common"
	"tinytiktok/user/models"
	"tinytiktok/user/proto/server"
	"tinytiktok/utils/config"
	"tinytiktok/utils/msg"
)

var UserDb *gorm.DB

type Handle struct {
	server.UnimplementedUserServiceServer
}

func init() {
	// 初始化配置文件
	path := os.Getenv("APP")
	dbConfig := config.NewConfig(fmt.Sprintf("%s/config", path), "mysql.yaml", "yaml")
	// 读取配置
	address := dbConfig.ReadString("User.Address")
	port := dbConfig.ReadInt("User.Port")
	username := dbConfig.ReadString("User.Username")
	password := dbConfig.ReadString("User.Password")
	dbName := dbConfig.ReadString("User.Dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbName)
	var err error
	UserDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if common.ServerMode == "release" {
		UserDb.Logger = logger.Default.LogMode(logger.Silent)
	}
	if err != nil {
		panic(msg.UnableConnectDB)
	}
	// 生成表
	Migrate()
}

func Migrate() {
	err := UserDb.AutoMigrate(&models.User{})
	if err != nil {
		panic(msg.UnableCreateTable)
	}
	err = UserDb.AutoMigrate(&models.Relation{})
	if err != nil {
		panic(msg.UnableCreateTable)
	}
	err = UserDb.AutoMigrate(&models.Message{})
	if err != nil {
		panic(msg.UnableCreateTable)
	}
}
