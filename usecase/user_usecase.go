package usecase

import (
	"blog_api/domain"
	"blog_api/infrastructure"
	"errors"
)

type UserUsecase interface {
	CreateUser(domain.User) error
	Login(string, string) (domain.User, error, string)
	FindUserById(string) (domain.User, error)
	FindAllUser() ([]domain.User, error)
	UpdateUser(domain.User) error
	DeleteUser(string) error
	FollowUser(string, string) error
}

type userUsecase struct {
	userRepo domain.UserRepository
	passwordService infrastructure.PasswordService
	jwtService infrastructure.JWTService
}

func NewUserUsecase(ur domain.UserRepository , ps infrastructure.PasswordService , jwtService infrastructure.JWTService) UserUsecase {
	return &userUsecase{
		userRepo: ur,
		passwordService: ps,
		jwtService: jwtService,
	}
}

func (u *userUsecase) CreateUser(user domain.User) error {
	if user.Username == "" {
		return errors.New("Invalid User ID")
	}
	if user.Username == "" {
		return errors.New("Invalid Username")
	}
	if user.Password == "" {
		return errors.New("Invalid Password")
	}

	newPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	
	user.Password = newPassword
	
	err = u.userRepo.CreateUser(user)
	if err != nil {
		return errors.New("Failed to create user")
	}
	return nil
}

func (u *userUsecase) Login(username string , password string) (domain.User, error , string) {
	user, err := u.userRepo.FindByUsername(username)
	if err != nil {
		return domain.User{}, errors.New("User not found") , ""
	}
	
	err = u.passwordService.ComparePassword(user.Password, password)
	if err != nil {
		return domain.User{}, errors.New("Invalid password") , ""
	}

	token, err := u.jwtService.GenerateToken(user)
	if err != nil {
		return domain.User{}, errors.New("Failed to generate token") , ""
	}

	return user, nil , token
}

func (u *userUsecase) FindUserById(id string) (domain.User, error) {
	// TODO: Implement FindUserById logic
	return domain.User{}, nil
}

func (u *userUsecase) FindAllUser() ([]domain.User, error) {
	// TODO: Implement FindAllUser logic
	return []domain.User{}, nil
}

func (u *userUsecase) UpdateUser(user domain.User) error {
	// TODO: Implement UpdateUser logic
	return nil
}

func (u *userUsecase) DeleteUser(id string) error {
	// TODO: Implement DeleteUser logic
	return nil
}

func (u *userUsecase) FollowUser(followerID string, followeeID string) error {
	// TODO: Implement FollowUser logic
	return nil
}

