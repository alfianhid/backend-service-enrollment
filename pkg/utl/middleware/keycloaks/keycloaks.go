// Package jsonwebtoken contains logic for using JSON web tokens
package keycloaks

import (
	"context"

	"backend-service/pkg/utl/config"

	"github.com/Nerzal/gocloak/v8"
)

// Service provides a Json-Web-Token authentication implementation
type Service struct {
	server        string
	realm         string
	client_id     string
	client_secret string
	user_admin    string
	pass_admin    string
	realm_admin   string
	client        gocloak.GoCloak
	ctx           context.Context
}

// New generates new JWT service necessery for auth middleware
func New(cfg *config.Keycloaks) *Service {
	return &Service{
		server:        cfg.Server,
		realm:         cfg.Realm,
		client_id:     cfg.ClientId,
		client_secret: cfg.ClientSecret,
		user_admin:    cfg.UserAdmin,
		pass_admin:    cfg.PassAdmin,
		realm_admin:   cfg.RealmAdmin,
		client:        gocloak.NewClient(cfg.Server),
		ctx:           context.Background(),
	}
}

func (j *Service) CreateUser(user gocloak.User) (*gocloak.JWT, string, error) {
	token, err := j.client.LoginAdmin(j.ctx, j.user_admin, j.pass_admin, j.realm_admin)
	if err != nil {
		return nil, "", err
	}

	userID, err := j.client.CreateUser(j.ctx, token.AccessToken, j.realm, user)
	if err != nil {
		return nil, "", err
	}

	return token, userID, nil
}

func (j *Service) SetUserPassword(token *gocloak.JWT, UserID, password string) error {
	err := j.client.SetPassword(j.ctx, token.AccessToken, UserID, j.realm, password, false)
	if err != nil {
		return err
	}

	return nil
}

// func (j *Service) SendVerifyEmail(token *gocloak.JWT, UserID string) error {
// 	params := gocloak.ExecuteActionsEmail{
// 		ClientID: &(j.client_id),
// 		UserID:   &UserID,
// 		Actions:  &[]string{"VERIFY_EMAIL"},
// 	}

// 	err := j.client.ExecuteActionsEmail(context.Background(), token.AccessToken, j.realm, params)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
