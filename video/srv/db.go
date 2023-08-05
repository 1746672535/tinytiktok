package srv

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"tinytiktok/utils/config"
	"tinytiktok/video/models"
	"tinytiktok/video/proto/server"
)

var VideoDb *gorm.DB

type Handle struct {
	server.UnimplementedVideoServiceServer
}

func init() {
	// 初始化配置文件
	path := os.Getenv("APP")
	dbConfig := config.NewConfig(fmt.Sprintf("%s\\config", path), "mysql.yaml", "yaml")
	// 读取配置
	address := dbConfig.ReadString("Video.Address")
	port := dbConfig.ReadInt("Video.Port")
	username := dbConfig.ReadString("Video.Username")
	password := dbConfig.ReadString("Video.Password")
	dbName := dbConfig.ReadString("Video.Dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, address, port, dbName)
	var err error
	VideoDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败: " + err.Error())
	}
	// 生成表
	Migrate()
}

func Migrate() {
	err := VideoDb.AutoMigrate(&models.Video{})
	if err != nil {
		panic("无法创建或迁移表")
	}
	err = VideoDb.AutoMigrate(&models.Like{})
	if err != nil {
		panic("无法创建或迁移表")
	}
}
