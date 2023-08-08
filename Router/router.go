package Router

import (
	"fmt"
	"taskup/Controller"
	"taskup/Middleware"

	"github.com/gin-gonic/gin"
)

func ServeApps() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	{
		AuthRoutes(authRoutes)
	}

	userRoutes := router.Group("/user")
	{
		UserRoutes(userRoutes)
	}

	projectRoutes := router.Group("/project")
	{
		ProjectRoutes(projectRoutes)
	}

	taskRoutes := router.Group("/task")
	{
		TaskRoutes(taskRoutes)
	}

	router.Run(":8080")
	fmt.Println("Server is running on port 8080")
}

func AuthRoutes(router *gin.RouterGroup) {
	router.POST("/register", Controller.Register)
	router.POST("/login", Controller.Login)
}

func UserRoutes(router *gin.RouterGroup) {
	//add protection middleware
	router.POST("/change-data", Middleware.UserMiddleware(), Controller.ChangePassword)
}

func ProjectRoutes(router *gin.RouterGroup) {
	router.POST("/create-project", Middleware.UserMiddleware(), Controller.CreateProject)
	router.GET("/get-project", Middleware.UserMiddleware(), Controller.GetProjects)
	router.GET("/get-project/:id", Middleware.UserMiddleware(), Controller.GetProjectByID)
	router.PUT("/update-project/:id", Middleware.UserMiddleware(), Controller.UpdateProject)
	router.DELETE("/delete-project/:id", Middleware.UserMiddleware(), Controller.DeleteProject)
}

func TaskRoutes(router *gin.RouterGroup) {
	router.POST("/create-task", Middleware.UserMiddleware(), Controller.CreateTask)
	router.GET("/get-task-by-project/:id_project", Middleware.UserMiddleware(), Controller.GetTasksByIdProject)
	router.GET("/get-task/:id", Middleware.UserMiddleware(), Controller.GetTasks)
	router.PUT("/update-task/:id", Middleware.UserMiddleware(), Controller.UpdateTask)
	router.DELETE("/delete-task/:id", Middleware.UserMiddleware(), Controller.DeleteTask)
}
