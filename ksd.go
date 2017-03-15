package main

import (
	"fmt"
	"os"
	"net/http"
	"ksd/controller"
)

func main() {
	var KSDListenPort = os.Getenv("KSD_PORT")
	if len(KSDListenPort) < 1 {
		KSDListenPort = "8080"
	}

	http.HandleFunc("/deploy", controller.DeployAction)
	err := http.ListenAndServe(fmt.Sprintf(":%s", KSDListenPort), nil)
	if nil != err {
		panic(err)
	}
}
