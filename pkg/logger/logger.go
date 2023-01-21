package logger

import (
	"errors"
	"log"
	"os"
	"time"
)

type Logger struct {
	fullPath   string
	file       *os.File
	dateFormat string
}

func NewLogger(fileName string, path string) (*Logger, error) {

	if !validateLoggerParams(&fileName, &path) {
		return nil, errors.New("can not create logger. File name or path is empty")
	}

	result := new(Logger)
	result.dateFormat = "2006-01-02 15:04:05"
	result.fullPath = path + fileName + " " + time.Now().Format(result.dateFormat) + ".log"

	err := createLogFile(&result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (l *Logger) WriteInfo(msg string) {
	logMsg := "INFO: " + msg + "\n"

	_, err := l.file.Write([]byte(time.Now().Format(l.dateFormat) + " " + logMsg))

	if err != nil {
		log.Println("Can not write info logs message: " + msg)
	}
}

func (l *Logger) WriteWarning(msg string) {
	logMsg := "WARNING: " + msg + "\n"

	_, err := l.file.WriteString(time.Now().Format(l.dateFormat) + " " + logMsg)

	if err != nil {
		log.Println("Can not write warning logs message: " + msg)
	}
}
func (l *Logger) WriteError(msg string) {
	logMsg := "ERROR: " + msg + "\n"

	_, err := l.file.WriteString(time.Now().Format(l.dateFormat) + " " + logMsg)

	if err != nil {
		log.Println("Can not write error logs message: " + msg)
	}
}

func createLogFile(logger **Logger) error {
	file, err := os.Create((*logger).fullPath)

	if err != nil {
		return err
	}

	(*logger).file = file
	return nil
}
func validateLoggerParams(fileName *string, path *string) bool {
	return *fileName != "" || *path != ""
}
