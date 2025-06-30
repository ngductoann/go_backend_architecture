package main

import "github.com/ngductoann/go_backend_architecture/internal/routers"

func main() {
	r := routers.NewRouter()
	r.Run() // listen and serve on ":8080" (Default)
}
