package mongo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock data for testing
var mockUsers = []UserInfo{
	{Id: "user1", Pw: "Password1", Name: "유저1"},
	{Id: "user2", Pw: "", Name: "유저2"},
	{Id: "user3", Pw: " ", Name: "유저3"},
}

var mockLoginSuccessUsers = []UserInfo{
	{Id: "user1", Pw: "Password1", Name: "유저1"},
	{Id: "user2", Pw: "", Name: "유저2"},
	{Id: "user3", Pw: " ", Name: "유저3"},
}

var mockLoginFailUsers = []UserInfo{
	{Id: "user1", Pw: "Password2", Name: "유저1"},
	{Id: "user1", Pw: "", Name: "유저1"},
	{Id: "user2", Pw: " ", Name: "유저2"},
	{Id: "user3", Pw: "", Name: "유저3"},
	{Id: "user4", Pw: "", Name: "유저4"},
}

func mockUserRepository() (UserRepository, error) {
	return NewUserRepository("user", "userInfoTest")
}

func TestClearCollection(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	err = repo.clearCollection()
	assert.Nil(t, err)
}

func TestRegister(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	for _, user := range mockUsers {
		err = Register(user.Id, user.Pw, user.Name, repo)
		assert.Nil(t, err)
	}
}

func TestRegister_Fail(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	for _, user := range mockUsers {
		err = Register(user.Id, user.Pw, user.Name, repo)
		var userAlreadyExistsError *UserAlreadyExistsError
		if errors.As(err, &userAlreadyExistsError) {
			assert.NotNil(t, userAlreadyExistsError, "Expected UserAlreadyExistsError, got nil")
		} else {
			t.Errorf("Expected UserAlreadyExistsError, got different error or nil")
		}
	}
}

func TestLogin_Success(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	for _, user := range mockLoginSuccessUsers {
		err = Login(user.Id, user.Pw, repo)
		assert.Nil(t, err)
	}
}

func TestLogin_Fail(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	for _, user := range mockLoginFailUsers {
		err = Login(user.Id, user.Pw, repo)
		assert.Error(t, err)
	}
}

func TestDBUserRepository_Close(t *testing.T) {
	repo, err := mockUserRepository()
	assert.Nil(t, err)
	err = repo.Close()
	assert.Nil(t, err)
}
