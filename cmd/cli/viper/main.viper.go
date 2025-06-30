package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/") // path to config
	viper.SetConfigName("local")      // name file
	viper.SetConfigType("yaml")       // file type

	// read configuration
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	// read
	fmt.Println("Server Port::", viper.GetInt("server.port"))          // read port server
	fmt.Println("Security key::", viper.GetString("security.jwt.key")) // read security key

	// configure structure
	var config Config
	if err := viper.Unmarshal(&config); err != nil { // unmarshal to struct
		fmt.Println("Error unmarshalling config:", err)
	}

	fmt.Println("Server Port from struct::", config.Server.Port) // read port server from struct

	for _, db := range config.Databases {
		fmt.Printf("Database User: %s, Password: %s, Host: %s\n", db.User, db.Password, db.Host)
	}
}
