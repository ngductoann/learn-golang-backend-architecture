package main

import (
	"fmt"

	"github.com/spf13/viper"
)

// multi config use struc
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`

	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Printf("Server Port:: ", viper.GetInt("server.port"))

	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode configuration %v \n", err)
	}

	fmt.Println("Config Port:: ", viper.GetInt("config.server.port"))
	fmt.Println("Config Port:: ", config.Server.Port)

	for _, db := range config.Databases {
		fmt.Printf("database User: %s, password: %s, host: %s \n", db.User, db.Password, db.Host)
	}
}
