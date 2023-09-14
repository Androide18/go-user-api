package services

import (
	"github.com/androide18/go-user-api/pkg/models"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{DB: db}
}

func (s *UserService) CreateUser(user *models.User) *models.User {
	s.DB.Create(user)
	return user
}

func (s *UserService) GetAllUsers() []models.User {
	var users []models.User
	s.DB.Find(&users)
	return users
}

func (s *UserService) GetUserById(id int64) (*models.User, error) {
	var user models.User
	err := s.DB.Where("ID=?", id).First(&user).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetUserByIdWithChannel(id int64) <-chan *models.User {
	userChan := make(chan *models.User)

	go func() {
		defer close(userChan)

		userDetails, err := s.GetUserById(id)
		if err != nil {
			//handle the error and send an error message through the channel
			userChan <- nil
			return
		}
		userChan <- userDetails
	}()

	return userChan
}

func (s *UserService) DeleteUser(id int64) models.User {
	var user models.User
	s.DB.Where("ID=?", id).Delete(user)
	return user
}

func (s *UserService) UpdateUser(user *models.User) (*models.User, error) {
	err := s.DB.Save(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
