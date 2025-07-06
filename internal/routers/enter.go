package routers

import (
	"github.com/ngductoann/go_backend_architecture/internal/routers/manage"
	"github.com/ngductoann/go_backend_architecture/internal/routers/user"
)

type RouterGroup struct {
	User    user.UserRouterGroup
	Manager manage.ManageRouterGroup
}

var RouterGroupApp = new(RouterGroup)
