package error_logger

import "log"

type ErrorPrinter struct{}

func (t *ErrorPrinter) LogError(err error) {
	log.Println(err)
}
