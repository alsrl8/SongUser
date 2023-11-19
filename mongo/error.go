package mongo

import "fmt"

type cannotConnectToMongoCloudError struct {
}

func (e *cannotConnectToMongoCloudError) Error() string {
	return fmt.Sprintf("can't connect to mongo cloud")
}

type userNotFoundError struct {
	Id string
}

func (e *userNotFoundError) Error() string {
	return fmt.Sprintf("user with Id(%s) not found", e.Id)
}

type userAlreadyExistsError struct {
	Id string
}

func (e *userAlreadyExistsError) Error() string {
	return fmt.Sprintf("user with Id(%s) already exists", e.Id)
}

type passwordMismatchError struct {
	Id string
}

func (e *passwordMismatchError) Error() string {
	return fmt.Sprintf("user password is mismatched with user id(%s)", e.Id)
}
