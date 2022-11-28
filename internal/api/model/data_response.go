package model

type DataResponse struct {
	Links   DataResponseLinks    `json:"links"`
	Buckets []DataResponseBucket `json:"buckets"`
}

type DataResponseLinks struct {
	Variables string `json:"$variables"`
}

type DataResponseBucket struct {
	Name string             `json:"name"`
	Data []DataResponseData `json:"data"`
}

type DataResponseData struct {
	Arg   string  `json:"arg"`
	Value float64 `json:"value"`
}
