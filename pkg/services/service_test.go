package services_test

import (
	"testing"

	"github.com/androide18/go-user-api/pkg/models"
	"github.com/androide18/go-user-api/pkg/services"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

// setup and cleanup for each test to ensure that tests are isolated and don't interfere with each other
func setupDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.DropTableIfExists(&models.User{}) // Clean up any previous data
	db.AutoMigrate(&models.User{})       // Create necessary tables
	return db
}
func TestGetAllUsers(t *testing.T) {
	db := setupDB()
	defer db.Close()

	userService := services.NewUserService(db)
	users := userService.GetAllUsers()

	assert.Equal(t, len(users), 0)
}

func TestCreateUser(t *testing.T) {
	db := setupDB()
	defer db.Close()

	userService := services.NewUserService(db)
	user := userService.CreateUser(&models.User{
		Name:     "Test",
		Lastname: "Test",
		Email:    "test@test.com",
	})

	assert.Equal(t, user.Name, "Test")
	assert.Equal(t, user.Lastname, "Test")
	assert.Equal(t, user.Email, "test@test.com")
}

func TestGetUserById(t *testing.T) {
	db := setupDB()
	defer db.Close()

	userService := services.NewUserService(db)
	user := userService.CreateUser(&models.User{
		Name:     "Test",
		Lastname: "Test",
		Email:    "test@test.com",
	})

	userDetails, err := userService.GetUserById(1)
	assert.Nil(t, err)
	assert.Equal(t, userDetails.Name, user.Name)
	assert.Equal(t, userDetails.Lastname, user.Lastname)
	assert.Equal(t, userDetails.Email, user.Email)
}

func TestUpdateUser(t *testing.T) {
	db := setupDB()
	defer db.Close()

	userService := services.NewUserService(db)
	user := userService.CreateUser(&models.User{
		Name:     "Test",
		Lastname: "Test",
		Email:    "test@test.com",
	})

	assert.Equal(t, user.Name, "Test")

	updatedUser := &models.User{
		Name:     "UpdatedName",
		Lastname: "UpdatedLastname",
		Email:    "updated@test.com",
	}

	updatedUser, err := userService.UpdateUser(updatedUser)
	assert.Nil(t, err)
	assert.Equal(t, updatedUser.Name, "UpdatedName")
	assert.Equal(t, updatedUser.Lastname, "UpdatedLastname")
	assert.Equal(t, updatedUser.Email, "updated@test.com")
}

func TestDeleteUser(t *testing.T) {
	db := setupDB()
	defer db.Close()

	userService := services.NewUserService(db)
	user := userService.CreateUser(&models.User{
		Name:     "Test",
		Lastname: "Test",
		Email:    "test@test.com",
	})

	assert.Equal(t, user.Name, "Test")

	userService.DeleteUser(1)

	// assert there is no users
	users := userService.GetAllUsers()
	assert.Equal(t, len(users), 0)
}
