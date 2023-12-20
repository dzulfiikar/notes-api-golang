package repositories

import (
	"fmt"
	schema "notes-api-golang/framework/sql/schemas"

	"gorm.io/gorm"
)

type UserRepository struct {
	dbInstance *gorm.DB
}

func NewUserRepository(dbInstance *gorm.DB) *UserRepository {
	return &UserRepository{
		dbInstance: dbInstance,
	}
}

func (repository *UserRepository) Save(userData schema.User) (user schema.User) {
	repository.dbInstance.Create(&userData)
	return userData
}

func (repository *UserRepository) FetchAll() (user schema.User) {
	repository.dbInstance.Find(&user)
	return user
}

func (repository *UserRepository) FetchUserExistsByEmail(email string) bool {
	return repository.dbInstance.Where("email = ?", email).Find(&schema.User{}).RowsAffected > 0
}

func (repository *UserRepository) FetchUserByEmail(email string) (user schema.User, err error) {
	repository.dbInstance.Where("email = ?", email).First(&user)

	if user.Email == "" {
		return user, fmt.Errorf("email not found")
	}

	return
}
