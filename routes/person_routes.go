package routes

import (
	"main/controller"

	"github.com/gin-gonic/gin"
)

func PersonRoutes(personController *controller.PersonController) *gin.Engine{
    router:=gin.Default()

    personRoutes:=router.Group("api/person")
    personRoutes.GET("",personController.FindAll)
    personRoutes.GET("/:personId",personController.FindById)
    personRoutes.POST("",personController.Create)
    personRoutes.PUT("/:personId",personController.Update)
    personRoutes.DELETE("/:personId",personController.Delete)

    return router

}