package api

type DataResponse struct {
	Links DataResponseLinks `json:"links"`
}

type DataResponseLinks struct {
	Variables string `json:"$variables"`
}
