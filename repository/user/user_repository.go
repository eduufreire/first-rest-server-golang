package user

import (
	"database/sql"
	"fmt"

	"github.com/eduufreire/rest-api-users/model"
	db "github.com/eduufreire/rest-api-users/repository"
)

type UserRepository struct {
	database *sql.DB
}

func New() *UserRepository {
	ur := UserRepository{}
	ur.database = db.Connect()
	return &ur
}

func (ur *UserRepository) CreateUser(user *model.User) (int, error) {
	stmt, err := ur.database.Prepare("insert into user(name, age, birthday) values (?, ?, ?)")
	if err != nil {
		return -1, fmt.Errorf("error in database")
	}

	result, err := stmt.Exec(user.Name, user.Age, user.Birthday)
	if err != nil {
		return -1, fmt.Errorf("error in database")
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("error on insert new user")
	}
	return int(lastInsertId), nil
}

func (ur *UserRepository) GetUserById(id int) *model.User {
	stmt, err := ur.database.Prepare("select * from user where id = ?")
	if err != nil {
		fmt.Print(err)
	}

	result, err := stmt.Query(id)
	if err != nil {
		fmt.Print(err)
	}

	user := model.User{}
	if result.Next() {
		result.Scan(&user.ID, &user.Name, &user.Age, &user.Birthday)
	}
	return &user
}

func (ur *UserRepository) GetAllUsers() *[]model.User {
	stmt, err := ur.database.Prepare("select * from user")
	if err != nil {
		fmt.Print(err)
	}

	result, err := stmt.Query()
	if err != nil {
		fmt.Print(err)
	}

	var users []model.User
	for result.Next() {
		var user model.User
		result.Scan(&user.ID, &user.Name, &user.Age, &user.Birthday)
		users = append(users, user)
	}
	return &users
}
