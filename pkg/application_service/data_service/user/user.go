// Package store contains the components necessary for api services
// to interact with the database
package user

import (
	"net/http"
	"strings"

	models "backend-service/pkg/utl/models"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

var (
	ErrAlreadyExists  = echo.NewHTTPError(http.StatusBadRequest, "email already exists")
	ErrRecordNotFound = echo.NewHTTPError(http.StatusNotFound, "email not found")
)

type UserDBClient struct{}

func NewUserDBClient() *UserDBClient {
	return &UserDBClient{}
}

func (u *UserDBClient) Create(db *gorm.DB, user models.Users) (*models.Users, error) {
	var checkUser = new(models.Users)

	// email check
	if err := db.Where(
		"lower(email) = ?",
		strings.ToLower(user.Email)).First(&checkUser).Error; err == nil {
		return nil, ErrAlreadyExists
	} else if !gorm.IsRecordNotFoundError(err) {
		return nil, err
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	// password masking
	user.Password = ""

	return &user, nil
}

func (u *UserDBClient) Update(db *gorm.DB, user *models.Users) error {
	return db.Save(user).Error
}
