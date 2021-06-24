package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/skinnykaen/mqtt-broker"
	"github.com/skinnykaen/mqtt-broker/package/handler"
	"github.com/skinnykaen/mqtt-broker/package/repository"
	"github.com/skinnykaen/mqtt-broker/package/service"
	"github.com/spf13/viper"
	"os"
)

func main() {
	//Формат логов консоли
	logrus.SetFormatter(new(logrus.JSONFormatter))

	//Инициализация файла конфигурации
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}
	//Загрузка файла конфигурации
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error initializing env variables: %s", err.Error())
	}

	//Инициализация базы данных
	db, err := repository.NewMysqlDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
	})
	if err != nil{
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//Внедрение зависимостей
	repos := repository.NewRepository(db)
	service := service.NewService(repos)
	handlers := handler.NewHandler(service)

	//Инициализация сервера
	srv := new(mqtt.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err!= nil {
		logrus.Fatalf("error occured while running server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}