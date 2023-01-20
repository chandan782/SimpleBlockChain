package transaction

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/rohanraj7316/logger"
	sLogger "gorm.io/gorm/logger"
)

const (
	logStrPrefix = "SQLLogger"
)

type Sqllogger struct {
	LogTrace bool
}

func (l Sqllogger) LogMode(level sLogger.LogLevel) sLogger.Interface {
	newLogger := l
	return &newLogger
}

func (l Sqllogger) Info(ctx context.Context, msg string, i ...interface{}) {
	fields := []logger.Field{
		{
			Key:   "sqlMessage",
			Value: msg,
		},
	}
	for key, value := range i {
		field := logger.Field{
			Key:   strconv.Itoa(key),
			Value: value,
		}
		fields = append(fields, field)
	}
	logger.Info(msg, fields...)
}

func (l Sqllogger) Warn(ctx context.Context, msg string, i ...interface{}) {
	fields := []logger.Field{
		{
			Key:   "sqlMessage",
			Value: msg,
		},
	}
	for key, value := range i {
		field := logger.Field{
			Key:   strconv.Itoa(key),
			Value: value,
		}
		fields = append(fields, field)
	}
	logger.Warn(msg, fields...)
}

func (l Sqllogger) Error(ctx context.Context, msg string, i ...interface{}) {
	fields := []logger.Field{
		{
			Key:   "sqlMessage",
			Value: msg,
		},
	}
	for key, value := range i {
		field := logger.Field{
			Key:   strconv.Itoa(key),
			Value: value,
		}
		fields = append(fields, field)
	}
	logger.Error(msg, fields...)
}

func (l Sqllogger) Trace(ctx context.Context, start time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogTrace {
		latency := time.Since(start).Round(time.Millisecond).String()
		query, rowAffected := fc()

		lStr := fmt.Sprintf("%s | %s | %s", logStrPrefix, latency, strconv.Itoa(int(rowAffected)))
		fmt.Printf("testing requestId: %s", ctx.Value("requestId"))
		fields := []logger.Field{
			{
				Key:   "requestId",
				Value: ctx.Value("requestId"),
			},
			{
				Key:   "latency",
				Value: latency,
			},
			{
				Key:   "query",
				Value: query,
			},
			{
				Key:   "rowAffected",
				Value: rowAffected,
			},
		}

		if err != nil {
			fields = append(fields, logger.Field{
				Key:   "error",
				Value: err.Error(),
			})
			logger.Error(lStr, fields...)
		} else {
			logger.Info(lStr, fields...)
		}
	}
}
