package logger_test

import (
	"testing"

	"github.com/kgoins/ldsview/logger"
)

type Result struct {
	jobID int
	txt   string
}

func Test(t *testing.T) {
	conf := logger.LoggerConfig{
		LogLevel:    string(logger.Debug),
		LogFilePath: "",
	}
	log := logger.NewZapLogger(conf)

	result := Result{
		jobID: 2,
		txt:   "udp_chargen_19",
	}

	log.Debug(
		"Job complete %d %s",
		result.jobID,
		result.txt,
	)
}
