package initialize

import (
	"fmt"

	"github.com/ngductoann/go_backend_architecture/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
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
	if err := viper.Unmarshal(&global.Config); err != nil { // unmarshal to struct
		fmt.Println("Error unmarshalling config:", err)
	}
}
