package error_logger

import "log"

type ErrorPrinter struct{}

func (t *ErrorPrinter) Log(err error) {
	log.Println(err)
}
