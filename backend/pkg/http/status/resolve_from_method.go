package status

import "net/http"

func ResolveFromMethod(req *http.Request) int {
	switch req.Method {
	case http.MethodPost:
		return http.StatusCreated
	}
	return http.StatusOK
}
