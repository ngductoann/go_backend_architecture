package initialize

import (
	"fmt"

	"github.com/ngductoann/go_backend_architecture/global"
)

func Run() {
	// load configuration
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql", m.Username, m.Password)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8002")
}
