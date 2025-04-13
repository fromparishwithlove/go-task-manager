package controllers

import (
	"net/http"
	"hello/config"
	"hello/models"

	"github.com/gin-gonic/gin"
)

// CreateTask — создаёт задачу
func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// тут можно вытаскивать ID пользователя из токена, пока подставим вручную
	username, _ := c.Get("username")

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	task.UserID = user.ID


	if err := config.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать задачу"})
		return
	}

	c.JSON(http.StatusOK, task)
}

// GetTasks — список всех задач
func GetTasks(c *gin.Context) {
	username, _ := c.Get("username")

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	var tasks []models.Task
	config.DB.Where("user_id = ?", user.ID).Find(&tasks)

	c.JSON(http.StatusOK, tasks)

}

// UpdateTask — обновляет задачу по id
func UpdateTask(c *gin.Context) {
	id := c.Param("id")

	username, _ := c.Get("username")

	// Найдём юзера по username
	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", id, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	// Обновляем из тела запроса
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&task)

	c.JSON(http.StatusOK, task)
}

// DeleteTask — удаляет задачу по id
func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	username, _ := c.Get("username")

	var user models.User
	if err := config.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не найден"})
		return
	}

	var task models.Task
	if err := config.DB.Where("id = ? AND user_id = ?", id, user.ID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Задача не найдена"})
		return
	}

	config.DB.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message": "Задача удалена"})
}
