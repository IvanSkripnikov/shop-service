package api

import (
	v1 "loyalty_system/api/v1"

	"net/http"
)

var Routes = map[string]func(http.ResponseWriter, *http.Request){
	"/health": v1.Health,
}
