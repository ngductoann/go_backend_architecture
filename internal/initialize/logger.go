package initialize

import (
	"github.com/ngductoann/go_backend_architecture/global"
	"github.com/ngductoann/go_backend_architecture/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
