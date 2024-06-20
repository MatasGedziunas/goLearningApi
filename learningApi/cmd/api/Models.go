package models

type Request struct {
	Url     string
	Headers map[string]string
}

type Response[T any] struct {
	Data       T
	StatusCode int
}

type Account struct {
	Balance  int
	Username string
}
