package dao

import (
	"database/sql"
	"go_dome/cache"
	"go_dome/models"
)

func CreateUser(user *models.User) error {
	_, err := cache.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	return err
}

func UpdateUser(user *models.User) error {
	_, err := cache.DB.Exec("UPDATE users SET username = ?, password = ? WHERE id = ?", user.Username, user.Password, user.ID)
	return err
}

func DeleteUser(id int) error {
	_, err := cache.DB.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}

func GetUserById(id int) (*models.User, error) {
	user := &models.User{}
	row := cache.DB.QueryRow("SELECT id, username, password FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
