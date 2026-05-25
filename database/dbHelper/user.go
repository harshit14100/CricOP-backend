package dbHelper

import (
	"context"

	"backend/database"
	"backend/models"
)

func CreateUser(user models.Users) error {

	query := `
	INSERT INTO users(name, phone_no, password)
	VALUES($1, $2, $3)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		user.Name,
		user.PhoneNo,
		user.Password,
	)

	return err
}

func GetUserByPhone(phone string) (*models.Users, error) {

	user := models.Users{}
	query := `
	SELECT id, name, phone_no, password
	FROM users
	WHERE phone_no = $1
	`

	row := database.DB.QueryRow(
		context.Background(),
		query,
		phone,
	)

	err := row.Scan(
		&user.ID,
		&user.Name,
		&user.PhoneNo,
		&user.Password,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func ResetPassword(phoneNo, hashedPassword string) error {
	query := `
UPDATE users
SET password = $1
WHERE phone_no = $2
`
	_, err := database.DB.Exec(
		context.Background(),
		query,
		phoneNo,
		hashedPassword,
	)
	return err
}

func GetAllUsers() ([]models.UserResponse, error) {

	query := `
	SELECT id,
	       name,
	       phone_no
	FROM users
	`

	rows, err := database.DB.Query(
		context.Background(),
		query,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.UserResponse

	for rows.Next() {

		var user models.UserResponse

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.PhoneNo,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
