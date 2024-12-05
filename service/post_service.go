package service

import (
	"go_dome/dao"
	"go_dome/models"
)

func CreatePost(post *models.Post) error {
	return dao.CreatePost(post)
}

func GetPostById(id int) ([]models.Post, error) {
	return dao.GetPostById(id)
}
