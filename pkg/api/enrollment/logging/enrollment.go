package enrollment

import (
	"time"

	"backend-service/pkg/api/enrollment"
	models "backend-service/pkg/utl/models"

	"github.com/labstack/echo"
)

const packageName = "enrollment"

// LogService represents user logging service
type LogService struct {
	enrollment.Service
	logger models.Logger
}

// New creates new user logging service
func New(svc enrollment.Service, logger models.Logger) *LogService {
	return &LogService{
		Service: svc,
		logger:  logger,
	}
}

// Create logging
func (ls *LogService) Register(c echo.Context, req models.User) (resp *models.RegisterResponse, err error) {
	dupe := req
	dupe.Password = "xxx-xxx-xxx"
	defer func(begin time.Time) {
		ls.logger.Log(
			c,
			packageName, "Register user request", err,
			map[string]interface{}{
				"req":  dupe,
				"resp": resp,
				"took": time.Since(begin),
			},
		)
	}(time.Now())
	return ls.Service.Register(c, req)
}
