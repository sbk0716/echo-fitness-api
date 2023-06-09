package repositories

import (
	"fitness-api/cmd/models"
	"fitness-api/cmd/storage"
	"fmt"
	"time"
)

func CreateUser(user models.User) (models.User, error) {
	fmt.Println("Creating user...")
	now := time.Now()
	db := storage.GetDB()
	sqlStatement := `
		INSERT INTO private.users (name, email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, name, email, password, created_at, updated_at`
	id := 0
	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password, now, now).Scan(&id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	fmt.Printf("Created user: %#v\n", user)
	user.Id = id
	fmt.Printf("Created user: %#v\n", user)
	return user, nil
}

func UpdateUser(user models.User, id int) (models.User, error) {
	now := time.Now()
	db := storage.GetDB()
	sqlStatement := `
		UPDATE private.users
		SET name = $2, email = $3, password = $4, updated_at = $5
		WHERE id = $1
		RETURNING id, name, email, password, created_at, updated_at`
	err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, now).Scan(&id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return models.User{}, err
	}
	fmt.Printf("Updated user: %#v\n", user)
	user.Id = id
	fmt.Printf("Updated user: %#v\n", user)

	return user, nil
}
