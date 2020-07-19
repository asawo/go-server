package db

import (
	"fmt"
	"testing"
)

type MockPostgresDb interface {
	GetUsers() []User
	CreateUser(name string)
	UpdateUser(id int, name string)
	DeleteUser(id int)
}

type Mock struct{}

var mock Mock

func (m Mock) GetUsers() []User {
	testUser := User{5, "Test"}
	testUser2 := User{6, "Arthur"}
	mockResponse := []User{
		testUser,
		testUser2,
	}
	return mockResponse
}

func (m Mock) CreateUser(name string) {
	testUser := User{5, name}
	Users = append(Users, testUser)
}

func (m Mock) UpdateUser(id int, name string) {

	for i, value := range Users {
		fmt.Println(i, value)
		if value.ID == id {
			Users[i] = User{id, name}
		}
	}
	fmt.Printf("User with id of %v does not exist", id)
}

func (m Mock) DeleteUser(id int) {
	testUser := User{5, "Test"}

	for i, value := range Users {
		fmt.Println(i, value)
		if value == testUser {
			Users = append(Users[:i], Users[i+1:]...)
		}
	}

	fmt.Println("Users", Users)
}

var pg PostgresDb = ConnectToDb()

func TestGetUsers(t *testing.T) {
	Users := pg.GetUsers()
	testUser := User{5, "Test"}

	if Users[1] != testUser {
		t.Errorf("User 1 should be %v, got %v", testUser, Users[1])
	}
}
