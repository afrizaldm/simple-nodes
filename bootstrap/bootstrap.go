package bootstrap

import (
	"nodes_app/router"

	"github.com/gin-gonic/gin"
)

type IBootstrap interface {
	Run(port ...string)
	SetupRouter(router *router.Router)
}

type Bootstrap struct {
	*gin.Engine
}

func New() *Bootstrap {
	return &Bootstrap{gin.New()}

}

func (b *Bootstrap) Run(port ...string) {
	p := ":8080"
	if len(port) > 0 {
		p = port[0]
	}

	b.Engine.Run(p)
}

func (b *Bootstrap) SetupRouter(router *router.Router) {
	e := router.SetupRouter(b.Engine)
	b.Engine = e
}
