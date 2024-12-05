package dao

import (
	"go_dome/cache"
	"go_dome/models"
)

func CreatePost(post *models.Post) error {
	_, err := cache.DB.Exec("INSERT INTO post (title, content,user_id,) VALUES (?, ?,?)", post.Title, post.Content, post.UserID)
	return err
}
func GetPostById(id int) ([]models.Post, error) {
	query := "SELECT id, title, content, user_id FROM post WHERE user_id = ?"
	rows, err := cache.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.UserID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func UpdatePost(post *models.Post) error {
	_, err := cache.DB.Exec("UPDATE post SET title = ?, content = ? WHERE id = ?", post.Title, post.Content, post.ID)
	return err
}

func DeletePost(id int) error {
	_, err := cache.DB.Exec("DELETE FROM post WHERE id = ?", id)
	return err
}
