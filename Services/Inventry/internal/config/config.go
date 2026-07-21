package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-require:"true"` // this are the struct tag that will tell to the cleanenv that what will be the name in the yml or env
	DbUri      string `yaml:"db_uri" env:"DB_URI" env-require:"true"`
	HttpServer `yaml:"http_server"`
}

func MustLoad() Config {

	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == ""{
	
		flags:= flag.String("inventory-config","","inventiory config")
		flag.Parse()
		configPath= *flags

		if configPath ==""{
			log.Fatal("Config path not provide")
		}
	}

	if _ ,err := os.Stat(configPath) ; os.IsNotExist(err){
			log.Fatal("Config file not exist !!!")
	}


	var cfg Config

	err := cleanenv.ReadConfig(configPath,&cfg)

	if err!= nil {
		 log.Fatalf("Error to read the config %s",err)
	}

	return  cfg

}