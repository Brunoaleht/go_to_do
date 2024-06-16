package configs

import (
	"log"

	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func init() {
	viper.SetDefault("api.port", 9000)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
}

func LoadConfig() error {
	// Carregar variáveis do arquivo .env
	err := gotenv.Load();
	if err != nil {
		log.Printf("No .env file found: %v", err)
	} else {
		log.Printf("Loaded .env file")
	}

	// Associar variáveis de ambiente com Viper
	viper.BindEnv("api.port", "API_PORT")
	viper.BindEnv("database.host", "DATABASE_HOST")
	viper.BindEnv("database.port", "DATABASE_PORT")
	viper.BindEnv("database.user", "DATABASE_USER")
	viper.BindEnv("database.password", "DATABASE_PASSWORD")
	viper.BindEnv("database.dbname", "DATABASE_NAME")

	// Carregar a configuração na struct

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Password: viper.GetString("database.password"),
		DBName:   viper.GetString("database.dbname"),
	}

	// Outra maneira de Carregar a configuração na struct
	// cfg = &config{
	// 	API: APIConfig{
	// 		Port: viper.GetString("api.port"),
	// 	},
	// 	DB: DBConfig{
	// 		Host:     viper.GetString("database.host"),
	// 		Port:     viper.GetString("database.port"),
	// 		User:     viper.GetString("database.user"),
	// 		Password: viper.GetString("database.password"),
	// 		DBName:   viper.GetString("database.dbname"),
	// 	},
	// }

	return nil
}

func GetDBConfig() DBConfig {
	return cfg.DB
}

func GetAPIConfig() APIConfig {
	return cfg.API
}

func GetServerPort() string {
	return cfg.API.Port
}