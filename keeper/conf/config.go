package conf

import (
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ConfigStruct struct {
	ChainID    int    `yaml:"chainID"`
	RPC        string `yaml:"rpc"`
	PrivateKey string `yaml:"privateKey"`
	Subgraph   string `yaml:"subgraph"`
	PriceFeed  string `yaml:"priceFeed"`
}

var Config ConfigStruct

func Init() {
	log.Println("Init config ....")
	viper := viper.New()
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	viper.Unmarshal(&Config)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err = viper.ReadInConfig()
		if err == nil {
			viper.Unmarshal(&Config)
			log.Printf("config: %+v", Config)
		} else {
			log.Printf("ReadInConfig error: %s", err)
		}
	})
}
