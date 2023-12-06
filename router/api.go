package router

import (
	"github.com/andrifirmansyah12/projectGo/controllers"
	"github.com/andrifirmansyah12/projectGo/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome To This Website")
	})
	api := r.Group("/api")
	{
		public := api.Group("/users")
		{
			public.POST("/login", controllers.Login)
			public.POST("/register", controllers.Signup)
		}
		protected := api.Group("/auth").Use(middlewares.Authz())
		{
			protected.PUT("/users/:userid", controllers.UpdateUsers)
			protected.DELETE("/users/:userid", controllers.DeleteUsers)

			protected.GET("/profile", controllers.Profile)
			protected.GET("/photos", controllers.FindPhotos)
			protected.POST("/photos", controllers.CreatePhoto)
			protected.GET("/photos/:id", controllers.FindPhoto)
			protected.PUT("/photos/:id", controllers.UpdatePhoto)
			protected.DELETE("/photos/:id", controllers.DeletePhoto)

		}
	}
	return r
}
