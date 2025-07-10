package repositories

import (
	"go_auth/databases/models"
	"go_auth/handlers/http/payloads/request"
	pagination "go_auth/lib"
	"log"
	"strings"

	"gorm.io/gorm"
)

type UserRepository interface {
	ListUsers(query request.ListUserRequest) (users []models.Users, totalRow int64, err error)
	RegisterUser(data models.Users) (user models.Users, err error)
	FindUserByUsername(username string) (user models.Users, err error)
	// Logout(request request.LogoutRequest) (err error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

// ListUsers implements UserRepository.
func (r *userRepositoryImpl) ListUsers(query request.ListUserRequest) (users []models.Users, totalRow int64, err error) {
	db := r.db.Model(&models.Users{})

	if query.Search != nil && len(*query.Search) > 0 {
		search := strings.ToLower(*query.Search)
		db = db.Where(`LOWER(username) LIKE ? OR LOWER(firstname) LIKE ? OR LOWER(lastname) LIKE ?`, "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	db.Count(&totalRow)

	pagination := pagination.GetOffset(query.Page, query.Size)
	db = db.Limit(pagination.Size).Offset(pagination.Offset)

	if err = db.Find(&users).Error; err != nil {
		log.Printf("REPO Users -> ListUsers, err: %s", err)
	}

	return
}

func (r *userRepositoryImpl) RegisterUser(data models.Users) (user models.Users, err error) {
	err = r.db.Create(&data).Error
	if err != nil {
		log.Printf("REPO Users -> RegisterUser, err: %s", err)
	}
	return
}

func (r *userRepositoryImpl) FindUserByUsername(username string) (user models.Users, err error) {
	err = r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("REPO Users -> LoginUser, err: %s", err)
	}
	return
}
