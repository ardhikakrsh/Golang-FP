package main

import (
	"fmt"
	"strings"

	"github.com/biskitsx/Leave-Management-System/config"
	"github.com/biskitsx/Leave-Management-System/repository"
	"github.com/biskitsx/Leave-Management-System/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	initConfig()
	app := initServer()
	db := initDatabase()
	router.Init(app, db)

	app.Listen(":8000")
}

func initServer() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: config.NewCustomConfigFiber,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000, http://client:3000",
		// AllowMethods:     strings.Split(viper.GetString("cors.allow_methods"), ","),
		// AllowHeaders:     strings.Split(viper.GetString("cors.allow_headers"), ","),
		AllowCredentials: true,
	}))
	return app
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	// config ค่าผ่าน cli
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v TimeZone=Asia/Bangkok",
		viper.GetString("db.host"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.database"),
		viper.GetString("db.port"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&repository.User{}, &repository.Leave{})
	if err != nil {
		panic(err)
	}
	return db
}
