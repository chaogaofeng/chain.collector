package logger

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
	tmlog "github.com/tendermint/tendermint/libs/log"
	"os"
	"path/filepath"
	"time"
)

var _ tmlog.Logger = (*ZeroLogWrapper)(nil)

// ZeroLogWrapper provides a wrapper around a zerolog.Logger instance. It implements
// Tendermint's Logger interface.
type ZeroLogWrapper struct {
	zerolog.Logger
}

// Info implements Tendermint's Logger interface and logs with level INFO. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Info(msg string, keyVals ...interface{}) {
	z.Logger.Info().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Error implements Tendermint's Logger interface and logs with level ERR. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Error(msg string, keyVals ...interface{}) {
	z.Logger.Error().Fields(getLogFields(keyVals...)).Msg(msg)
}

// Debug implements Tendermint's Logger interface and logs with level DEBUG. A set
// of key/value tuples may be provided to add context to the log. The number of
// tuples must be even and the key of the tuple must be a string.
func (z ZeroLogWrapper) Debug(msg string, keyVals ...interface{}) {
	z.Logger.Debug().Fields(getLogFields(keyVals...)).Msg(msg)
}

// With returns a new wrapped logger with additional context provided by a set
// of key/value tuples. The number of tuples must be even and the key of the
// tuple must be a string.
func (z ZeroLogWrapper) With(keyVals ...interface{}) tmlog.Logger {
	return ZeroLogWrapper{z.Logger.With().Fields(getLogFields(keyVals...)).Logger()}
}

func getLogFields(keyVals ...interface{}) map[string]interface{} {
	if len(keyVals)%2 != 0 {
		return nil
	}

	fields := make(map[string]interface{})
	for i := 0; i < len(keyVals); i += 2 {
		fields[keyVals[i].(string)] = keyVals[i+1]
	}

	return fields
}

func init() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05.999"
}

func NewLogger(level string, dir string) ZeroLogWrapper {
	logWriter := zerolog.ConsoleWriter{Out: os.Stderr}
	if len(dir) > 0 {
		rl, err := rotatelogs.New(filepath.Join(dir, "logs", "log%Y%m%d"),
			rotatelogs.WithRotationTime(time.Hour*24), // 日志切割时间间隔
			rotatelogs.WithMaxAge(time.Hour*24*30),    //文件存活时间
		)
		if err != nil {
			panic(fmt.Errorf("failed to parse log file: %w", err))
		}
		logWriter = zerolog.ConsoleWriter{Out: rl}
	}

	logLvl, err := zerolog.ParseLevel(level)
	if err != nil {
		panic(fmt.Errorf("failed to parse log level (%s): %w", viper.GetString("cyy.log.level"), err))
	}
	return ZeroLogWrapper{zerolog.New(logWriter).Level(logLvl).With().Timestamp().Logger()}
}
