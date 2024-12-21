package v1

import (
	"loyalty_system/apiresponse"

	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	data := apiresponse.ResponseData{
		"status": "OK2",
	}

	apiresponse.SendResponse(w, data, "response")
}
