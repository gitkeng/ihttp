package log_test

import (
	"github.com/gitkeng/ihttp/log"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestDebug(t *testing.T) {
	log.Debug("test debug")
}

func TestDebugf(t *testing.T) {
	log.Debugf("test debugf %s", "test")
}

func TestDebugj(t *testing.T) {
	json := log.JSON{"test": "test"}
	log.Debugj("test debug json", "json", json)
}

func TestDebugWithNewFile(t *testing.T) {
	log.New(
		log.WithLevel(zapcore.DebugLevel),
		log.WithFileLocation("test.log"),
	)
	json := log.JSON{"test": 1}
	log.Debugj("test debug json", "json", json)
}
