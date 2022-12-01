package handler

import "fmt"

func createApiUrl(endpoint string, params ...any) string {
	return "http://localhost:3030" + fmt.Sprintf(endpoint, params...)
}
