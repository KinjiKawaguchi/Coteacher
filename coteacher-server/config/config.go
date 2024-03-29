package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBUser      string `envconfig:"DB_USER_NAME" required:"true"`
	DBPassword  string `envconfig:"DB_PASSWORD" required:"true"`
	DBAddress   string `envconfig:"DB_ADDR" required:"true"`
	DBName      string `envconfig:"DB_NAME" required:"true"`
	FrontEndURL string `envconfig:"FRONTEND_URL" required:"true"`
	Port        string `envconfig:"PORT" required:"true"`
}

func New() (*Config, error) {
	// .env ファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
		return nil, err
	}

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		log.Printf("Error processing env variables: %v", err)
		return nil, err
	}

	// デバッグログを出力
	log.Println("DB User:", config.DBUser)
	log.Println("DB Address:", config.DBAddress)
	log.Println("DB Name:", config.DBName)
	log.Println("FrontEnd URL:", config.FrontEndURL)
	log.Println("Port:", config.Port)

	return config, nil
}
