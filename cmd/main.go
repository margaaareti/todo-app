package main

import (
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/pkg/handler"
	"github.com/zhashkevych/todo-app/pkg/repository"
	"github.com/zhashkevych/todo-app/pkg/service"
)

func main() {
	//Считывает значение config и записывает во внутренний файл viper
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err)
	}

	//Загружаем переменную окружения из .env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	//Инициализируем базу, передав все необходимые значения
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),     //"localhost",
		Port:     viper.GetString("db.port"),     //"5432",
		Username: viper.GetString("db.username"), //"postgres",
		Password: os.Getenv("DB_PASSWORD"),       //"qwerty",
		DBName:   viper.GetString("db.dbname"),   //"postgres",
		SSLMode:  viper.GetString("db.sslmode"),  //"disable",
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	//Объявляем зависимости в нужном порядке:

	//1.Создаем сначала репозиторий
	repos := repository.NewRepository(db)
	//2.Сервис, который зависит от repository
	services := service.NewService(repos)
	//3. Handler, который зависит от service
	handlers := handler.NewHandler(services)

	//Инициализируем экземпляр сервера
	srv := new(todo.Server)
	//Запуск сервера вызовом метода Run
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}

}

//Функция для инициализации конфигурационных данных
func initConfig() error {
	//Указываем имя директории
	viper.AddConfigPath("configs")
	//Указываем имя файла
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	//Считывает значение config и записывает во внутренний объект viper
	return viper.ReadInConfig()
}
