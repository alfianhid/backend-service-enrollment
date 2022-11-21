package api

import (
	"crypto/sha1"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	// pkg/api
	"backend-service/pkg/api/enrollment"
	el "backend-service/pkg/api/enrollment/logging"
	et "backend-service/pkg/api/enrollment/transport"

	// pkg/utl
	"backend-service/pkg/utl/config"
	"backend-service/pkg/utl/datastore"
	keycloakService "backend-service/pkg/utl/middleware/keycloaks"
	"backend-service/pkg/utl/models"
	"backend-service/pkg/utl/secure"
	"backend-service/pkg/utl/server"
	"backend-service/pkg/utl/zlog"
)

// newServices initializes new services for API
func newServices(cfg *config.Configuration) (sec *secure.Service, jwt *keycloakService.Service, log *zlog.Log, e *echo.Echo) {
	sec = secure.New(cfg.App.MinPasswordStr, sha1.New())
	jwt = keycloakService.New(cfg.Keycloaks)
	log = zlog.New()
	e = server.New()

	return sec, jwt, log, e
}

// initializeControllers initializes new HTTP services for each controller
func initializeControllers(db *gorm.DB, sec *secure.Service, jwt *keycloakService.Service, log *zlog.Log, e *echo.Echo) {
	api := e.Group("/api")

	et.NewHTTP(el.New(enrollment.Initialize(db, sec, jwt), log), api)
}

// startServer starts HTTP server with correct config & initialized services
func startServer(e *echo.Echo, cfg *config.Configuration) {
	server.Start(e, &server.Config{
		Port:                cfg.Server.Port,
		ReadTimeoutSeconds:  cfg.Server.ReadTimeout,
		WriteTimeoutSeconds: cfg.Server.WriteTimeout,
		Debug:               cfg.Server.Debug,
	})
}

// Start starts the API service
func Start(cfg *config.Configuration) error {
	db, err := datastore.NewMySQLGormDb(cfg.DB)
	if err != nil {
		return err
	}

	db.AutoMigrate(&models.Users{}) // migrating Users model to datbase table

	sec, jwt, log, e := newServices(cfg)

	initializeControllers(db, sec, jwt, log, e)

	e.Static("/swaggerui", cfg.App.SwaggerUIPath)

	startServer(e, cfg)

	return nil
}
