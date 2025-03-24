package global

import (
	"os"
	"strings"
)

var Mode = strings.ToLower(os.Getenv("MODE"))
var LogLevel = strings.ToLower(os.Getenv("LOG_LEVEL"))

var Version = "0.0.1"

var RouterPrefix = "api/v1"

var OriginalPassword = "123456"
