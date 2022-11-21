package enrollment

import (
	"backend-service/pkg/application_service/data_service/user"
	models "backend-service/pkg/utl/models"

	"github.com/Nerzal/gocloak/v8"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// Securer represents security interface
type Securer interface {
	Hash(string) string
	Password(string, ...string) bool
}

type DBClientInterface interface {
	Create(*gorm.DB, models.Users) (*models.Users, error)
	Update(*gorm.DB, *models.Users) error
}

type Keycloak interface {
	CreateUser(gocloak.User) (*gocloak.JWT, string, error)
	SetUserPassword(*gocloak.JWT, string, string) error
	// SendVerifyEmail(*gocloak.JWT, string) error
}

// Service represents user application interface
type Service interface {
	Register(echo.Context, models.User) (*models.RegisterResponse, error)
}

// RequestHandler represents user application service
type RequestHandler struct {
	sec   Securer
	cloak Keycloak
	db    *gorm.DB
	udb   DBClientInterface
}

// New creates new user RequestHandler application service
func New(db *gorm.DB, udb DBClientInterface, sec Securer, cloak Keycloak) *RequestHandler {
	return &RequestHandler{db: db, udb: udb, sec: sec, cloak: cloak}
}

// Initialize initalizes User RequestHandler application service with defaults
func Initialize(db *gorm.DB, sec Securer, cloak Keycloak) *RequestHandler {
	return New(db, user.NewUserDBClient(), sec, cloak)
}
