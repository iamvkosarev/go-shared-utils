package logs

import (
	"fmt"
	"net/http"
)

type HttpLogger struct {
	logErrorsIntoConsole  bool
	logSuccessIntoConsole bool
}

func NewHttpLogger(logSuccessIntoConsole bool, logErrorsIntoConsole bool) HttpLogger {
	return HttpLogger{logSuccessIntoConsole: logSuccessIntoConsole, logErrorsIntoConsole: logErrorsIntoConsole}
}

func (l HttpLogger) Error(writer http.ResponseWriter, message string, statusCode int) {
	if statusCode < 300 {
		fmt.Errorf("wrong status code: %d. expected error", statusCode)
	}
	http.Error(writer, message, statusCode)
	if l.logErrorsIntoConsole {
		fmt.Printf("(%v) %s\n", statusCode, message)
	}
}

func (l HttpLogger) InternalError(writer http.ResponseWriter, message string, statusCode int, err error) {
	if statusCode < 300 {
		fmt.Errorf("wrong status code: %d. expected error", statusCode)
	}
	http.Error(writer, message, statusCode)
	if l.logErrorsIntoConsole {
		fmt.Printf("(%v) %s : %s\n", statusCode, message, err.Error())
	}
}

func (l HttpLogger) Success(writer http.ResponseWriter, message string, statusCode int) {
	if statusCode < 200 || statusCode > 300 {
		fmt.Errorf("wrong status code: %d. expected success", statusCode)
	}
	writer.WriteHeader(statusCode)
	fmt.Fprint(writer, message)
	if l.logSuccessIntoConsole {
		fmt.Printf("(%v) %s\n", statusCode, message)
	}
}
