package api

type DataResponse struct {
	Links DataResponseLinks `json:"links"`
}

type DataResponseLinks struct {
	Subject string `json:"$subject"`
}
