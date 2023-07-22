package infra

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

func InitLogger() (*zap.Logger, error) {

	// TODO: need to add the appropriate logging location as well as the centralized logging setup
	logFilePath := "myLog.log"
	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer logFile.Close()

	configJson := []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "` + logFilePath + `"],
		"errorOutputPaths": ["stderr", "` + logFilePath + `"],
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"levelEncoder": "lowercase"
		  }
	}`)

	var zapConfig zap.Config
	if err := json.Unmarshal(configJson, &zapConfig); err != nil {
		panic(err)
	}

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Info("created and initialized the zap logger")
	return logger, nil
}
