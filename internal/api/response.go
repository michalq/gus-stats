package api

type ApiReponse[T any] struct {
	Data T `json:"data"`
}
