package users

type UsersRepository interface {
	RetrieveUser(userID string)(UserData, error)
	CreateUser(user *UserData) (int, error)
	UpdateUser(user *UserData) error
}