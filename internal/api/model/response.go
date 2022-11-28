package model

type ApiReponse[T any] struct {
	Data T `json:"data"`
}
