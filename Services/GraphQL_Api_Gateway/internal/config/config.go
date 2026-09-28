package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type GrpcClient struct {
	UserClient string `yaml:"user_client" env-required:"true"`
	// AddressClient  string `yaml:"address_client" env-required:"true"`
	// CartClient  string `yaml:"cart_client" env-required:"true"`
	// CheckoutClient  string `yaml:"checkout_client" env-required:"true"`
	// InvetoryClient  string `yaml:"inventory_client" env-required:"true"`
	// OrderClient  string `yaml:"order_client" env-required:"true"`
	ProductClient  string `yaml:"product_client" env-required:"true"`
	WarehouseClient  string `yaml:"warehouse_client" env-required:"true"`
}

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}


type Config struct {
	HttpServer `yaml:"http_server"`
	GrpcClient 	`yaml:"grpc_client"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == ""{
	flags := flag.String("gateway-config","","gateway config path ")
	flag.Parse()
	configPath = *flags

		if configPath == ""{
			log.Fatal("Config Path not provide ")
		}
	}

	_ ,err:=os.Stat(configPath)

	if os.IsNotExist(err) {
		log.Fatal("Config file path not exist check again !!!")
	}

	var cfg Config 

	err=cleanenv.ReadConfig(configPath,&cfg)
	if err!= nil {
		log.Fatal("Error to read the configfile ",err)
	}

	return &cfg 
}