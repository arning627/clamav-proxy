package web

import (
	"fmt"
	"net/http"

	"github.com/arning627/clamav-proxy/config"
	clamd "github.com/arning627/clamav-proxy/internal"
)

func Ping(response http.ResponseWriter, request *http.Request) {
	client := clamd.NewClient(config.InitializeConfig.Clamav.Host, config.InitializeConfig.Clamav.Port)
	res, e := client.Ping()
	if e != nil {
		fmt.Fprintln(response, e.Error())
		return
	}
	fmt.Fprintln(response, res)
}
