package db_test

import (
	"fmt"
	"go-server/server/db"
	_ "go-server/server/db"
	"testing"
)

type mock struct{}

var m mock

var Users []db.User
var postgresDb db.PostgresDb
var postgres db.Postgres

func (m mock) GetUsers() []db.User {
	testUser := db.User{5, "Test"}
	testUser2 := db.User{6, "Arthur"}
	mockResponse := []db.User{
		testUser,
		testUser2,
	}
	return mockResponse
}

func (m mock) CreateUser(name string) {
	testUser := db.User{5, name}
	Users = append(Users, testUser)
}

func (m mock) UpdateUser(id int, name string) {

	for i, value := range Users {
		fmt.Println(i, value)
		if value.ID == id {
			Users[i] = db.User{id, name}
		}
	}
	fmt.Printf("User with id of %v does not exist", id)
}

func (m mock) DeleteUser(id int) {
	testUser := db.User{5, "Test"}

	for i, value := range Users {
		fmt.Println(i, value)
		if value == testUser {
			Users = append(Users[:i], Users[i+1:]...)
		}
	}

	fmt.Println("Users", Users)
}

func TestGetUsers(t *testing.T) {
	// pg := db.ConnectToDb()
	// Users := pg.GetUsers()
	testUser := db.User{5, "Test"}
	Users := db.RandomFunction(m)
	fmt.Println(Users)
	if Users[0] != testUser {
		t.Errorf("User 1 should be %v, got %v", testUser, Users[1])
	}
}
