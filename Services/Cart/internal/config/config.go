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

type GrpcServ struct {
	Port string `yaml:"port" env-required:"true"`
}

type GrpcClient struct {
	ProductClient  string `yaml:"product_client" env-required:"true"`
}


type Config struct {
	DB string `yaml:"db"  env-required:"true"`
	HttpServer `yaml:"http_server"`
	GrpcServ 	`yaml:"grpc_serv"`
	GrpcClient 	`yaml:"grpc_client"`
}

func MustLoad() Config{
	var configPath string

	configPath=os.Getenv("CONFIG_PATH")

	if configPath ==""{
		flags := flag.String("cart-config","","config path")
		flag.Parse()
		configPath =*flags

		if configPath == ""{
			log.Fatal("No Config Path provide ")
		}
	}

	if _,err:= os.Stat(configPath) ; os.IsNotExist(err){
		log.Fatal("No file exist Provide path")
	}

	var cfg Config

	err:= cleanenv.ReadConfig(configPath,&cfg)

	if err!= nil {
		log.Fatal("Erorr to read the env",err)
	}

	return cfg

}