package app

import (
	"devEnv/src/config"
	"net/http"
)

func StartHTTP() error {
	r := Routes()
	return http.ListenAndServe(":"+config.PORT, r)

}
