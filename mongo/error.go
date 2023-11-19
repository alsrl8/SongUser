package mongo

import "fmt"

type CannotConnectToMongoCloudError struct {
}

func (e *CannotConnectToMongoCloudError) Error() string {
	return fmt.Sprintf("can't connect to mongo cloud")
}

type UserNotFoundError struct {
	Id string
}

func (e *UserNotFoundError) Error() string {
	return fmt.Sprintf("user with Id(%s) not found", e.Id)
}

type UserAlreadyExistsError struct {
	Id string
}

func (e *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with Id(%s) already exists", e.Id)
}

type PasswordMismatchError struct {
	Id string
}

func (e *PasswordMismatchError) Error() string {
	return fmt.Sprintf("user password is mismatched with user id(%s)", e.Id)
}
