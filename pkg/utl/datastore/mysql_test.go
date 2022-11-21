package datastore_test

import (
	"backend-service/pkg/utl/config"
	"backend-service/pkg/utl/datastore"
	models "backend-service/pkg/utl/models"
	"backend-service/pkg/utl/support"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatDSN(t *testing.T) {
	cfgPath := support.TestingConfigPath()
	cfg, err := config.LoadConfigFrom(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if cfg == nil {
		t.Fatal(errors.New("unknown error loading testing yaml file"))
	}
	expectedDSN := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Protocol,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.Settings,
	)
	dsn := datastore.FormatDSN(cfg.DB)
	assert.Equal(t, expectedDSN, dsn, "dsn should be properly formatted")
}

func TestNew(t *testing.T) {
	cfgPath := support.DevelopmentConfigPath()
	cfg, err := config.LoadConfigFrom(cfgPath)
	if err != nil {
		t.Fatal(err)
	}
	if cfg == nil {
		t.Fatal(errors.New("unknown error loading testing yaml file"))
	}

	dsn := datastore.FormatDSN(cfg.DB)
	expectedDsn := "root:Admin1234$#@!" +
		fmt.Sprintf("@tcp(localhost:%s)/store_microservice", cfg.DB.Port) +
		"?tls=skip-verify&charset=utf8&parseTime=True&loc=Local&autocommit=true&timeout=20s"
	assert.Equal(t, expectedDsn, dsn, "dsn should be properly formated")

	corruptedDBcfg := &config.Database{
		Dialect:  cfg.DB.Dialect,
		User:     cfg.DB.User,
		Password: cfg.DB.Password,
		Name:     cfg.DB.Name,
		Protocol: cfg.DB.Protocol,
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Settings: cfg.DB.Settings,
	}
	corruptedDBcfg.Host, corruptedDBcfg.Port = "pluto", "53456345634563"
	_, err = datastore.NewMySQLGormDb(corruptedDBcfg)
	assert.EqualError(t, err, err.Error(), "there should be an error connecting to mysql with bad config")

	corruptedDBcfg.Host, corruptedDBcfg.Port = cfg.DB.Host, cfg.DB.Port
	corruptedDBcfg.Password = "root"
	_, err = datastore.NewMySQLGormDb(corruptedDBcfg)
	assert.EqualError(t, err, err.Error(), "there should be an error connecting to mysql with bad config")

	corruptedDBcfg.Password = "admin"
	_, err = datastore.NewMySQLGormDb(corruptedDBcfg)
	assert.EqualError(t, err, err.Error(), "there should be an error connecting to mysql with bad config")

	db, err := datastore.NewMySQLGormDb(cfg.DB)
	if err != nil {
		t.Fatalf("Error establishing connection %v", err)
	}
	modelsList := []interface{}{
		&models.Users{},
	}
	for _, model := range modelsList {
		if db.HasTable(model) {
			if err := db.DropTable(model).Error; err != nil {
				t.Fatalf("Error establishing connection %v", err)
			}
		}
		if err := db.CreateTable(model).Error; err != nil {
			t.Fatalf("Error establishing connection %v", err)
		}
	}
	assert.Nil(t, db.Close(), "there should not be an error closing the DB")
	db.Close()
}
