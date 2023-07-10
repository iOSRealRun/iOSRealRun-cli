package logger

import (
	"fmt"
	"log"
	"os"
)

type LogLevel int

const (
	Panic LogLevel = iota
	Fatal
	Error
	Warn
	Info
	Debug
	Trace
)

var LevelMap = map[string]LogLevel{
	"panic": Panic,
	"fatal": Fatal,
	"error": Error,
	"warn":  Warn,
	"info":  Info,
	"debug": Debug,
	"trace": Trace,
}

type myLogger struct {
	logger *log.Logger
	level  LogLevel
}

func NewMyLogger(filename string, log_level string) *myLogger {
	var level LogLevel
	if _, err := os.Stat("./log"); os.IsNotExist(err) {
		os.Mkdir("./log", 0777)
	}

	file, err := os.OpenFile(fmt.Sprintf("./log/%s", filename), os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}

	level, ok := LevelMap[log_level]
	if !ok {
		log.Fatal("log-level error")
	}

	return &myLogger{
		logger: log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile),
		level:  level,
	}
}

func (l *myLogger) Panic(v ...any) {
	if l.level >= Panic {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Panicf(format string, v ...any) {
	if l.level >= Panic {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Panicln(v ...any) {
	if l.level >= Panic {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Fatal(v ...any) {
	if l.level >= Fatal {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Fatalf(format string, v ...any) {
	if l.level >= Fatal {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Fatalln(v ...any) {
	if l.level >= Fatal {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Error(v ...any) {
	if l.level >= Error {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Errorf(format string, v ...any) {
	if l.level >= Error {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Errorln(v ...any) {
	if l.level >= Error {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Warn(v ...any) {
	if l.level >= Warn {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Warnf(format string, v ...any) {
	if l.level >= Warn {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Warnln(v ...any) {
	if l.level >= Warn {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Info(v ...any) {
	if l.level >= Info {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Infof(format string, v ...any) {
	if l.level >= Info {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Infoln(v ...any) {
	if l.level >= Info {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Debug(v ...any) {
	if l.level >= Debug {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Debugf(format string, v ...any) {
	if l.level >= Debug {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Debugln(v ...any) {
	if l.level >= Debug {
		l.logger.Println(v...)
	}
}

func (l *myLogger) Trace(v ...any) {
	if l.level >= Trace {
		l.logger.Print(v...)
	}
}
func (l *myLogger) Tracef(format string, v ...any) {
	if l.level >= Trace {
		l.logger.Printf(format, v...)
	}
}
func (l *myLogger) Traceln(v ...any) {
	if l.level >= Trace {
		l.logger.Println(v...)
	}
}
