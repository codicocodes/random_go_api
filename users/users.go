package users

import (
	"database/sql"
	"log"
	"time"
)

type UserRepository interface {
	List() []User
}

type userPostgreSQL struct {
	db *sql.DB
}

func GetUserService(db *sql.DB) userPostgreSQL {
	return userPostgreSQL{db: db}
}

type User struct {
	ID        int64     `field:"id" json:"id"`
	Username  string    `field:"username" json:"username"`
	CreatedAt time.Time `field:"createdAt" json:"createdAt"`
}

func (s userPostgreSQL) List() []User {
	rows, err := s.db.Query("SELECT id, username, \"createdAt\" FROM \"User\";")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		u := new(User)
		err = rows.Scan(&u.ID, &u.Username, &u.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, *u)
	}
	return users
}
