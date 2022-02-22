package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

func Log() *logrus.Logger {
	return log
}

func PrettyPrint(data interface{}) {
	res, _ := json.MarshalIndent(data, "", "    ")
	fmt.Printf("%v\n", string(res))
}
