package repo

type UserRepo struct{}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (up *UserRepo) GetInfoUser() string {
	return "ngductoann"
}
