package db

import (
	"testing"
)

// type SQLDB interface {
// 	Exec(query, args string) (sql.Result, error)
// }

// type MockDB struct {
// 	callParams []interface{}
// }

// func (mdb *MockDB) Exec(query, args string) (sql.Result, error) {
// 	mdb.callParams = []interface{}{query}
// 	mdb.callParams = append(mdb.callParams, args)

// 	return nil, nil
// }

// func (mdb *MockDB) CalledWith() []interface{} {
// 	return mdb.callParams
// }

var pg PostgresDb = ConnectToDb()

func TestGetUsers(t *testing.T) {
	Users := pg.GetUsers()
	testUser := User{5, "Test"}

	if Users[1] != testUser {
		t.Errorf("User 1 should be {5 Test}, got %v", Users[1])
	}
}
