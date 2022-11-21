// Package transport contains the HTTP service for user interactions
package transport

import (
	"net/http"

	"backend-service/pkg/api/enrollment"
	models "backend-service/pkg/utl/models"

	"github.com/labstack/echo"
)

// Custom errors
var (
	ErrUnknownPayload      = echo.NewHTTPError(http.StatusBadRequest, "payload is unknown")
	ErrPasswordsNotMaching = echo.NewHTTPError(http.StatusBadRequest, "password do not match")
)

// HTTP represents user http service
type HTTP struct {
	svc enrollment.Service
}

// NewHTTP creates new user http service
func NewHTTP(svc enrollment.Service, er *echo.Group) {
	h := HTTP{svc}

	er.POST("/register", h.register)
}

// createReq is a used to serialize the request payload to a struct
type createReq struct {
	Email            string `json:"email" validate:"required,email"`
	Username         string `json:"username" validate:"required,min=3,alphanum"`
	FirstName        string `json:"first_name" validate:"required,min=3"`
	LastName         string `json:"last_name" validate:"alpha"`
	Phone            string `json:"phone" validate:"required,min=6,max=15,numeric"`
	Company          string `json:"company" validate:"required,min=3"`
	BusinessRelation string `json:"business_relation" validate:"required,min=3"`
	Password         string `json:"password" validate:"required,min=8"`
	PasswordConfirm  string `json:"password_confirm" validate:"required,min=8"`
}

// register Creates new user account
//
// usage: POST /v1/users users userCreate
//
// responses:
//  200: userResp
//  400: errMsg
//  401: err
//  403: errMsg
//  500: err
func (h *HTTP) register(c echo.Context) error {
	r := new(createReq)

	if err := c.Bind(r); err != nil {
		return ErrUnknownPayload
	}

	if r.Password != r.PasswordConfirm {
		return ErrPasswordsNotMaching
	}

	usr, err := h.svc.Register(c, models.User{
		Email:            r.Email,
		Username:         r.Username,
		FirstName:        r.FirstName,
		LastName:         r.LastName,
		Phone:            r.Phone,
		Company:          r.Company,
		BusinessRelation: r.BusinessRelation,
		Password:         r.Password,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, usr)
}
