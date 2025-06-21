package usecases

import "taskManager/internal/domain"

type UserUseCase struct {
	repo domain.UserRepo
}

func NewUserUseCase(repo domain.UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}
func (u *UserUseCase) CreateUser(user domain.User) error {
	return u.repo.Create(user)
}
func (u *UserUseCase) UpdateUser(user domain.User) error {
	return u.repo.Update(user)
}
func (u *UserUseCase) DeleteUser(user domain.User) error {
	return u.repo.Delete(user)
}
func (u *UserUseCase) FindByUsername(un string) (domain.User, error) {
	return u.repo.FindByUsername(un)
}
