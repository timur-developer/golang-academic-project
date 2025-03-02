package userService

type UserService struct {
	repo UserRepository
}

func NewService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *UserService) UpdateUserByID(id uint, updates map[string]interface{}) (User, error) {
	return s.repo.UpdateUserByID(id, updates)
}

func (s *UserService) DeleteUserByID(id uint) (User, error) {
	return s.repo.DeleteUserByID(id)
}
