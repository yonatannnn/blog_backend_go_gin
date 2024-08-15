package usecase

import (
	"blog_api/domain"
	"blog_api/infrastructure"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserUsecase interface {
	CreateUser(domain.User) error
	Login(string, string) (domain.User, error, string)
	FindUserByUsername(string) (domain.User, error)
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
		return errors.New("Invalid UserName")
	}
	if user.Password == "" {
		return errors.New("Invalid Password")
	}


	newPassword, err := u.passwordService.HashPassword(user.Password)
	if err != nil {
		return errors.New("Failed to hash password")
	}
	
	user.Password = newPassword
	user.ID = primitive.NewObjectID()

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

func (u *userUsecase) FindUserByUsername(username string) (domain.User, error) {
	user , err := u.userRepo.FindByUsername(username)
	if err != nil {
		return domain.User{}, errors.New("User not found")
	}
	return user, nil
}

func (u *userUsecase) FindAllUser() ([]domain.User, error) {
	users , err := u.userRepo.FindAllUsers()
	if err != nil {
		return nil, errors.New("Failed to find users")
	}
	if len(users) > 0 {
		return users, nil
	}

	return users , nil
}

func (u *userUsecase) UpdateUser(user domain.User) error {
	// TODO: Implement UpdateUser logic
	return nil
}

func (u *userUsecase) DeleteUser(id string) error {
	
	objId , err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("Invalid ID")
	}
	err = u.userRepo.DeleteUser(objId)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) FollowUser(followerID string, followeeID string) error {
	// TODO: Implement FollowUser logic
	return nil
}

