package controllers

import (
	"net/http"

	"github.com/DrewButNotBarrymore/go-task/models"
	"github.com/gin-gonic/gin"
)

type CreateTaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	UserID      uint   `json:"user_id" binding:"required"`
	Priv1       string `json:"priv1"`
	Priv2       string `json:"priv2"`
}

type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
}

type CreateStatusInput struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTaskInput struct {
	ID          uint   `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      uint   `json:"user_id"`
	StatusID    uint   `json:"status_id"`
	Priv1       string `json:"priv1"`
	Priv2       string `json:"priv2"`
}

// GET all tasks @ "/tasks" endpoint
func FindTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Select("id", "title", "description", "user_id", "status_id").Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GET "/tasks/user/:userid"
func FindUserTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Select("id", "title", "description", "user_id", "status_id").Where("user_id = ?", c.Param("userid")).Find(&tasks)
	if len(tasks) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": tasks})

}

// POST to "/tasks" to add new task
func CreateTask(c *gin.Context) {
	// validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	task := models.Task{Title: input.Title, Description: input.Description, UserID: input.UserID, StatusID: 1}
	models.DB.Create(&task)

	// create history
	history := models.History{TaskID: task.ID, Status: task.StatusID, UserID: task.UserID}
	models.DB.Create(&history)

	c.JSON(http.StatusOK, gin.H{"data": task})

}

// POST to "/users" to add new user
func CreateUser(c *gin.Context) {
	// validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create user
	user := models.User{Name: input.Name}
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// PATCH "/task/:id" to to update task
func UpdateTask(c *gin.Context) {
	// get model if it exists
	var task models.Task
	if err := models.DB.Select("id", "title", "description", "user_id", "status_id").Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&task).Updates(input)

	// create history
	history := models.History{TaskID: task.ID, Status: task.StatusID, UserID: task.UserID}
	models.DB.Create(&history)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE "/tasks/:id" to soft delete task(status id=5)
func DeleteTask(c *gin.Context) {
	// get model if it exists
	var task models.Task
	if err := models.DB.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// update task status in delete
	task.StatusID = 5
	models.DB.Save(&task)

	// create history
	history := models.History{TaskID: task.ID, Status: task.StatusID, UserID: task.UserID}
	models.DB.Create(&history)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
