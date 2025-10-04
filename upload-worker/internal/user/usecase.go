package user

type UserUseCase struct {
	Repository *UserRepository
}

func NewUserUseCase(Repository *UserRepository) *UserUseCase {
	return &UserUseCase{
		Repository: Repository,
	}
}
