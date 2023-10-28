package helpers

import (
	"Booking/internal/config"
	"fmt"
	"net/http"
	"runtime/debug"
)

var appConfig *config.AppConfig

// NewHelpers set up app config for helper
func NewHelpers(a *config.AppConfig) {
	appConfig = a
}

func ClientError(writer http.ResponseWriter, status int) {
	appConfig.InfoLog.Println("client error with status of", status)
	http.Error(writer, http.StatusText(status), status)
}

func ServerError(writer http.ResponseWriter, err error) {
	trace := fmt.Sprint("%s\n%s", err.Error(), debug.Stack())
	appConfig.ErrorLog.Println(trace)
	http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
