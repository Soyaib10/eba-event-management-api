package models

import (

	"github.com/Soyaib10/eba-event-booking-api/db"
	"github.com/Soyaib10/eba-event-booking-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	// Insert statement and prepare it
	qurey := `INSERT INTO users(email, password)
			VALUES(?, ?)`
	stmt, err := db.DB.Prepare(qurey)
	if err != nil {
		return err
	}
	defer stmt.Close()
	
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	// Execute stmt
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	u.ID = userId
	return err
}