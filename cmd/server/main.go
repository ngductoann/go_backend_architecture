package main

import (
	"github.com/ngductoann/go_backend_architecture/internal/initialize"
)

func main() {
	// r := routers.NewRouter()
	//
	// // InitMysql()
	// // InitRedis()
	// // InitKafka()
	// // .....
	//
	// r.Run() // listen and serve on ":8080" (Default)

	initialize.Run()
}
