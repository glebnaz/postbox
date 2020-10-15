package logics

import (
	"errors"
	"fmt"

	"github.com/glebnaz/postbox/internal/entities"
)

const emptyUsersError = "empty users"

//GetUsers return users,find by ids
//if ids is empty get all users
func GetUsers(repository entities.UserRepository, request UserReq) ([]entities.User, error) {
	return repository.Get(request.IDs)
}

func InsertUsers(repository entities.UserRepository, request UserReq) error {
	//validate users
	if len(request.Users) == 0 {
		return errors.New(emptyUsersError)
	}
	for i, v := range request.Users {
		if len(v.SMTPAddress) == 0 {
			return fmt.Errorf("smtp_address in user %v empty", i+1)
		}
		if len(v.SMTPHost) == 0 {
			return fmt.Errorf("smtp_host in user %v empty", i+1)
		}
		if len(v.SMTPUser) == 0 {
			return fmt.Errorf("smtp_user in user %v empty", i+1)
		}
		if len(v.SMTPPass) == 0 {
			return fmt.Errorf("smtp_pass in user %v empty", i+1)
		}
		err := repository.Insert(v)
		if err != nil {
			return err
		}
	}
	return nil
}
