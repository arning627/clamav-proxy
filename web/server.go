package web

import (
	"fmt"
	"net/http"

	clamd "github.com/Arning627/clamav-proxy/internal"
)

func Ping(response http.ResponseWriter, request *http.Request) {

	client := clamd.NewClient("localhost", 3310)
	res, e := client.Ping()
	if e != nil {
		fmt.Fprintln(response, e.Error())
		return
	}
	fmt.Fprintln(response, res)
}
