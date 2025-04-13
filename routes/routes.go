package routes

import (
	"github.com/gin-gonic/gin"
	"hello/controllers"
	"hello/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login) // 🔐 Добавили маршрут для входа

	// Защищённый роут
	authorized := r.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware()) // подключаем проверку токена
	{
		authorized.GET("/me", controllers.Me) // теперь /me доступен только с токеном

		// 🔥 Задачи
		authorized.POST("/tasks", controllers.CreateTask)
		authorized.GET("/tasks", controllers.GetTasks)
		authorized.PUT("/tasks/:id", controllers.UpdateTask)
		authorized.DELETE("/tasks/:id", controllers.DeleteTask)
	}
}