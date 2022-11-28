package handler

import "fmt"

func createApiUrl(endpoint string, params ...any) string {
	return "http://localhost:3000" + fmt.Sprintf(endpoint, params...)
}
