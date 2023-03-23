package web

import (
	"fmt"
	"log"
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

func Scan(response http.ResponseWriter, request *http.Request) {

	client := clamd.NewClient(config.InitializeConfig.Clamav.Host, config.InitializeConfig.Clamav.Port)

	request.ParseForm()

	file, _, e := request.FormFile("file")

	if e != nil {
		log.Fatalln(e)
	}

	client.SacnStream(file)

}

func Execute(response http.ResponseWriter, request *http.Request) {
	client := clamd.NewClient(config.InitializeConfig.Clamav.Host, config.InitializeConfig.Clamav.Port)
	query := request.URL.Query()
	cmd, _ := query["cmd"]
	client.Execute(cmd[0])
	response.Write([]byte("ok..."))
}
