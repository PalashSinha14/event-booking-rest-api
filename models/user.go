package models

import (
	"errors"
	"github.com/palashsinha14/go-rest-api/db"
	"github.com/palashsinha14/go-rest-api/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"` 
}

func (u *User) Save() error {

	// PostgreSQL uses $1, $2 instead of ? placeholders
	query := `
	INSERT INTO users (email, password)
	VALUES ($1, $2)
	RETURNING id;
	`

	// Prepare statement using global DB connection
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	// Always close prepared statement after execution
	defer stmt.Close()

	// Hash password before storing
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// PostgreSQL does not support LastInsertId()
	// Instead we use QueryRow + Scan with RETURNING id
	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}


/*
func (u User) Save() error{
	query := "INSERT INTO users(email, password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err!=nil{
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
	result,err:= stmt.Exec(u.Email, hashedPassword)
	if err!=nil{
		return err
	}
	userId, err := result.LastInsertId()
	u.ID=userId
	return err
}


func (u *User) ValidateCredentials() error{
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil{
		return errors.New("Credentials invalid")
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid{
		return errors.New("Credentails invalid")
	}
	return nil
}
*/


func (u *User) ValidateCredentials() error {

	// PostgreSQL uses $1 instead of ?
	query := `
	SELECT id, password
	FROM users
	WHERE email = $1;
	`

	// Execute query
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	// Scan returned values into struct
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials invalid")
	}

	// Compare hashed password with entered password
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}