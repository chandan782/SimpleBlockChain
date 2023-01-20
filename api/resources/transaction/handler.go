package transaction

import (
	"fmt"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	EnvDBHost          = "POSTGRESQL_HOST"
	EnvDBUser          = "POSTGRESQL_USER"
	EnvDBPassword      = "POSTGRESQL_PASSWORD"
	EnvDBName          = "POSTGRESQL_DB_NAME"
	EnvDBLoggingEnable = "POSTGRESQL_LOG_TRACE"
	EnvDBPort          = "POSTGRESQL_PORT"
)

type TableType string

const (
	TRANSACTIONS TableType = "TRANSACTIONS"
)

type Handler struct {
	TableName string
	Db        *gorm.DB
}

func loadConfig() map[string]string {
	requiredFields := []string{
		EnvDBHost,
		EnvDBUser,
		EnvDBPassword,
		EnvDBName,
		EnvDBLoggingEnable,
		EnvDBPort,
	}
	fields := map[string]string{}
	for _, val := range requiredFields {
		fields[val] = os.Getenv(val)
	}

	return fields
}

func New(tType TableType) (*Handler, error) {
	envCfg := loadConfig()

	// initialize sql logger
	newLogger := Sqllogger{}
	logTrace := envCfg[EnvDBLoggingEnable]
	if logTrace == "true" {
		newLogger.LogTrace = true
	} else {
		newLogger.LogTrace = false
	}

	dsnStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		envCfg[EnvDBHost],
		envCfg[EnvDBUser],
		envCfg[EnvDBPassword],
		envCfg[EnvDBName],
		envCfg[EnvDBPort])
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: dsnStr}), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		return nil, err
	}

	tableName := strings.ToUpper(fmt.Sprintf("%s_TRANSACTIONS", string(tType)))

	return &Handler{
		Db:        db,
		TableName: tableName,
	}, nil
}
