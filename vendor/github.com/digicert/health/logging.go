package health

import (
	"fmt"
	"log"
)

func Panic(message string, params ...interface{}) {
	logCount.WithLabelValues("panic").Inc()
	if logLevel >=0 {
		log.Panicf("%s%s %s", "[PANIC]", debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Fatal(message string, params ...interface{}) {
	logCount.WithLabelValues("fatal").Inc()
	if logLevel >=1 {
		log.Fatalf("%s%s %s", "[FATAL]", debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Error(message string, params ...interface{}) {
	logCount.WithLabelValues("error").Inc()
	if logLevel >= 2 {
		log.Printf("%s%s %s","[ERROR]",debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Warn(message string, params ...interface{}) {
	logCount.WithLabelValues("warn").Inc()
	if logLevel >= 3 {
		log.Printf("%s%s %s","[WARN]",debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Info(message string, params ...interface{}) {
	logCount.WithLabelValues("info").Inc()
	if logLevel >= 4 {
		log.Printf("%s%s %s","[INFO]", debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Debug(message string, params ...interface{}) {
	logCount.WithLabelValues("debug").Inc()
	if logLevel >= 5 {
		log.Printf("%s%s %s","[DEBUG]", debugInfo(), fmt.Sprintf(message, params...))
	}
}

func Trace(message string, params ...interface{}) {
	logCount.WithLabelValues("trace").Inc()
	if logLevel >= 6 {
		log.Printf("%s%s %s","[TRACE]", debugInfo(), fmt.Sprintf(message, params...))
	}
}