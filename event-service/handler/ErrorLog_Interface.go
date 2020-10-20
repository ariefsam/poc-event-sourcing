package handler

type ErrorLog interface {
	Log(err error)
}
