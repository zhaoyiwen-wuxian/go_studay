package controllers

import (
	"go_dome/models"
	"go_dome/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.CreatePost(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post created successfully"})
}

func GetPost(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("user_id"))
	post, err := service.GetPostById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "post not found"})
		return
	}

	c.JSON(http.StatusOK, post)
}
