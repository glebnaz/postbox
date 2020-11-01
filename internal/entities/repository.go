package entities

const Collection = "postbox_users"

//UserRepository providing access to user data
type UserRepository interface {
	Get(ids []string) ([]User, error)
	Insert(object User) error
	Update(object User) error
	Delete(ids []string) error
}
