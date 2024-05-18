package models

import (
	"errors"

	"github.com/Milan-CS03/GO_REST/db"
	"github.com/Milan-CS03/GO_REST/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userid, err := result.LastInsertId()
	u.ID = userid
	return err
}

func (u User) ValidateCredentials() error {
	query := "SELECT password FROM users where email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&retrivedPassword)
	if err != nil {
		return errors.New("credential invalid")
	}
	if utils.CheckPasswordHash(u.Password, retrivedPassword) != true {
		return errors.New("credential invalid")
	}
	return nil
}
