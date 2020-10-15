package logics

import (
	"github.com/glebnaz/postbox/internal/entities"
)

//GetUsers return users,find by ids
//if ids is empty get all users
func GetUsers(repository entities.UserRepository, request UserReq) ([]entities.User, error) {
	return repository.Get(request.IDs)
}
