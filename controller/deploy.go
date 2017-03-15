package controller

import (
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"ksd/service"
	"fmt"
)

func DeployAction(w http.ResponseWriter, r *http.Request) {
	if "POST" != r.Method {
		setHttpStatus(w, http.StatusMethodNotAllowed)
		return
	}

	token := r.URL.Query().Get("token")
	if !isTokenValid(token) {
		setHttpStatus(w, http.StatusForbidden)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if nil != err || len(body) < 1 {
		setHttpStatus(w, http.StatusBadRequest)
		return
	}

	dpl := new(Deploy)
	err = json.Unmarshal(body, dpl)
	if nil != err {
		setHttpStatus(w, http.StatusBadRequest)
	}

	// TODO: check content-type

	clientset := service.NewKubernetesClientset()

	dp, err := clientset.Deployments(dpl.Namespace).Get(dpl.Deployment)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Deployment started for " + dpl.Namespace + "/" + dpl.Deployment)
	for idx,c := range dp.Spec.Template.Spec.Containers {
		for _, dpl_c := range dpl.Containers {
			if c.Name == dpl_c.Name {
				dp.Spec.Template.Spec.Containers[idx].Image = dpl_c.Image
				fmt.Println(dpl.Namespace + "/" + dpl.Deployment + "- Container " + dpl_c.Name + " Found, image has been set: " + dp.Spec.Template.Spec.Containers[idx].Image)
			}
		}
	}

	dp, err = clientset.Deployments(dpl.Namespace).Update(dp)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Deployment " + dpl.Namespace + "/" + dpl.Deployment + " updated")
}

func setHttpStatus(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)
}

func isTokenValid(token string) bool {
	validToken := os.Getenv("KSD_TOKEN")

	if len(validToken) < 1 {
		panic("ENV 'KSD_TOKEN' not set!")
	}

	return validToken == token
}

type Deploy struct {
	Namespace string `json:"namespace"`
	Deployment string `json:"deployment"`
	Containers []Container `json:"containers"`
}

type Container struct {
	Name string `json:"name"`
	Image string `json:"image"`
}
