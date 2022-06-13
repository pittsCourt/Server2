package health

import (
	"fmt"
	"runtime"
	"strings"
)

func SetDebug(in bool) {
	debug = in
}

func GetDebug() bool {
	return debug
}

func GetLogLevel() string {
	switch logLevel {
	case 0:
		return "panic"
	case 1:
		return "fatal"
	case 2:
		return "error"
	case 3:
		return "warn"
	case 4:
		return "info"
	case 5:
		return "debug"
	case 6:
		return "trace"
	default:
		return ""
	}
}

func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "panic":
		logLevel = 0
	case "fatal":
		logLevel = 1
	case "error":
		logLevel = 2
	case "warn":
		logLevel = 3
	case "info":
		logLevel = 4
	case "debug":
		logLevel = 5
	case "trace":
		logLevel = 6
	default:
		logLevel = 4
	}
}

func getNames(pc uintptr) (name string, packageName string) {
	name = runtime.FuncForPC(pc).Name()
	nameSplit := strings.Split(name, ".")
	name = nameSplit[len(nameSplit)-1]
	packageName = nameSplit[len(nameSplit)-2]
	packageNameParts := strings.Split(packageName, "/")
	packageName = strings.ToUpper(packageNameParts[len(packageNameParts)-1])
	return
}

func debugInfo() string {
	if debug {
		pc, _, line, ok := runtime.Caller(2)
		var name, packageName string
		if ok {
			name, packageName = getNames(pc)
		} else {
			name, packageName = "get_Debug_Info", "Error_In_Caller"
		}
		return fmt.Sprintf("[%s:%s:%d]", packageName, name, line)
	}
	return ""
}

// reduce the cardinality on paths
func parsePath(path string) string {
	if strings.Contains(path, "/rest/api/2/issue") {
		return "/rest/api/2/issue"
	} else if strings.Contains(path, "/api/v1/users/") {
		return "/api/v1/users"
	}

	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return "/"+parts[1]+"/"+parts[2]
	}

	return path
}
