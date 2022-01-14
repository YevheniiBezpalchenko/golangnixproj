package repositories

import (
	"database/sql"
	"github.com/YevheniiBezpalchenko/golangnixproj/internal/models"
	"log"
)

type UserDB struct {
	conn *sql.DB
}

func (u *UserDB) Construct(conn *sql.DB) {
	u.conn = conn
}

func (u *UserDB) GetByEmail(email string) models.User {
	var user models.User
	err := u.conn.QueryRow("select * from users where email = ?", email).Scan(&user.UserId, &user.Email, &user.PasswordHash)
	if err != nil {
		log.Fatal(err)
	}

	return user
}
func (u *UserDB) GetById(Id int) models.User {
	var user models.User
	err := u.conn.QueryRow("select * from users where id = ?", Id).Scan(&user.UserId, &user.Email, &user.PasswordHash)
	if err != nil {
		log.Fatal(err)
	}
	return user
}
func (u *UserDB) Create(user *models.User) (id int64, err error) {

	stmt, err := u.conn.Prepare("INSERT TO users(id,email,password) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(user.UserId, user.Email, user.PasswordHash)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	return lastId, err
}
func (u *UserDB) GetAll() (userArray []models.User) {
	var user models.User
	rows, err := u.conn.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&user.UserId, &user.Email, &user.PasswordHash)
		if err != nil {
			log.Fatal(err)
		}
		userArray = append(userArray, user)
	}
	return userArray
}
func (u *UserDB) Delete(Id int) error {
	_, err := u.conn.Query("DELETE * from users WHERE id=?", Id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}
func (u *UserDB) Update(id int, user *models.User) models.User {
	var updatedUser models.User = *user
	_, err := u.conn.Query("ALTER UPDATE users SET email=?, password=? WHERE id=?", user.Email, user.PasswordHash, id)
	if err != nil {
		log.Fatal(err)
	}
	return updatedUser
}
