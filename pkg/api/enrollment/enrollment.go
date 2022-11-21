package enrollment

import (
	"net/http"

	models "backend-service/pkg/utl/models"

	"github.com/Nerzal/gocloak/v8"

	"github.com/labstack/echo"
)

// Custom errors
var (
	ErrInsecurePassword = echo.NewHTTPError(http.StatusBadRequest, "insecure password")
)

func (a *RequestHandler) Register(c echo.Context, req models.User) (*models.RegisterResponse, error) {
	// check password strength
	if ok := a.sec.Password(req.Password, req.Email, req.Username, req.FirstName, req.LastName); !ok {
		return nil, ErrInsecurePassword
	}

	// register to Keycloak
	token, userID, err := a.cloak.CreateUser(gocloak.User{
		Email:     gocloak.StringP(req.Email),
		Username:  gocloak.StringP(req.Username),
		FirstName: gocloak.StringP(req.FirstName),
		LastName:  gocloak.StringP(req.LastName),
		Enabled:   gocloak.BoolP(true),
		// RequiredActions: &[]string{"VERIFY_EMAIL"},
	})
	if err != nil {
		return nil, err
	}

	// store user password to Keycloak
	err = a.cloak.SetUserPassword(token, userID, req.Password)
	if err != nil {
		return nil, err
	}

	// send email verification
	// err = a.cloak.SendVerifyEmail(token, userID)
	// if err != nil {
	// 	return nil, err
	// }

	// register to DB
	a.udb.Create(a.db, models.Users{
		UserID:           userID,
		Email:            req.Email,
		Username:         req.Username,
		FirstName:        req.FirstName,
		LastName:         req.LastName,
		Phone:            req.Phone,
		Company:          req.Company,
		BusinessRelation: req.BusinessRelation,
		Password:         a.sec.Hash(req.Password),
	})

	return &models.RegisterResponse{Status: 200, Message: userID}, nil
}
