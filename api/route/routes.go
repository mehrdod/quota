package route

import (
	"net/http"

	"alif/quota/app/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, apiVersion string) {

	vRouter := router.Group(apiVersion)
	{
		// Quota
		vRouter.POST("/quota", controllers.CreateQuota)
		vRouter.PATCH("/quota/id/:id", controllers.UpdateQuota)
		vRouter.DELETE("/quota/id/:id", controllers.RemoveQuota)
		vRouter.GET("/quota/all", controllers.GetAllQuotas)
		vRouter.GET("/quota/category/:category", controllers.GetQuotaByCategory)
		vRouter.GET("/quota/random", controllers.GetRandomQuota)

	}

	// Others
	router.GET("/ping", ping)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, `404`)
	})

}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
