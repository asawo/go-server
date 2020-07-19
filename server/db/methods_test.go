package db_test

import (
	"fmt"
	"go-server/server/db"
	_ "go-server/server/db"
	"testing"
)

type Mock struct{}

var mock Mock

var Users []db.User

func (m Mock) GetUsers() []db.User {
	testUser := db.User{5, "Test"}
	testUser2 := db.User{6, "Arthur"}
	mockResponse := []db.User{
		testUser,
		testUser2,
	}
	return mockResponse
}

func (m Mock) CreateUser(name string) {
	testUser := db.User{5, name}
	Users = append(Users, testUser)
}

func (m Mock) UpdateUser(id int, name string) {

	for i, value := range Users {
		fmt.Println(i, value)
		if value.ID == id {
			Users[i] = db.User{id, name}
		}
	}
	fmt.Printf("User with id of %v does not exist", id)
}

func (m Mock) DeleteUser(id int) {
	testUser := db.User{5, "Test"}

	for i, value := range Users {
		fmt.Println(i, value)
		if value == testUser {
			Users = append(Users[:i], Users[i+1:]...)
		}
	}

	fmt.Println("Users", Users)
}

var pg db.PostgresDb = db.ConnectToDb()

func TestGetUsers(t *testing.T) {
	Users := pg.GetUsers()
	testUser := db.User{5, "Test"}

	if Users[1] != testUser {
		t.Errorf("User 1 should be %v, got %v", testUser, Users[1])
	}
}
