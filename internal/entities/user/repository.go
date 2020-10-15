package user

const Collection = "postbox_users"

//Repository providing access to user data
type Repository interface {
	Get(ids ...string) ([]User, error)
	Insert(object User) error
	Update(object User) error
	Delete(id string) (User, error)
}
