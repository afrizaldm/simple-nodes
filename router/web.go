package router

import (
	"nodes_app/app/controllers/web"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	SetupRouter(app *gin.Engine) *gin.Engine
}

type Router struct{}

func New() *Router {
	return &Router{}
}

func (r *Router) SetupRouter(app *gin.Engine) *gin.Engine {

	FolderController := web.Node{}

	app.GET("/", web.Index)
	app.GET("/app", web.App)
	app.GET("/ping", web.Ping)

	app.GET("/nodes", FolderController.Index)
	app.GET("/nodes/:id", FolderController.Show)
	app.GET("/nodes/create", FolderController.Create)
	app.POST("/nodes", FolderController.Store)
	app.GET("/nodes/:id/edit", FolderController.Edit)
	app.PUT("/nodes/:id", FolderController.Update)
	app.PATCH("/nodes/:id", FolderController.Update)
	app.DELETE("/nodes/:id", FolderController.Delete)

	return app

}
