package global

import (
	"github.com/ngductoann/go_backend_architecture/pkg/logger"
	"github.com/ngductoann/go_backend_architecture/pkg/setting"
)

var (
	Config setting.Config    // set global variable for configuration
	Logger *logger.LoggerZap // set global variable for logger
)

/*
Config
Redis
Mysql
...
*/
