package slackapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"main/database"
	"main/jenkinsbuilds"
	"net/http"
	"os"
)

func slackWebhook(query string) {
	url := os.Getenv("WEBHOOK_URL")
	fmt.Println("URL:>", url)

	var jsonStr = []byte(query)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func SWBuildNaskStarted(username string, module string, version string, repository string) {

	slackWebhook(`{"text":"Poszedł build ` + module + ` o wersji ` + version + ` na repozytorium: ` + repository + `. Wykonany przez: ` + username + `"}`)
}

func SWBuildNaskEnded(module string, version string, repository string, branch string) {

	job := jenkinsbuilds.BuildJob{Module: module, Version: version, Repository: repository, Branch: branch}

	clientFirestore := database.Initialisation()
	database.UpdateBuild(clientFirestore, job)
	slackWebhook(`{"text":"Skończył się build ` + module + ` o wersji ` + version + ` na repozytorium: ` + repository + ` z brancza: ` + branch + `"}`)
}

func SWDeploy(username string, module string, version string, environment string) {

	slackWebhook(`{"text":"Wykonano deploy ` + module + ` o wersji ` + version + ` na środowisko: ` + environment + `. Wykonany przez: ` + username + `"}`)

}

func SWInsufficientRole(name string) {

	slackWebhook(`{"text":"Rola ` + name + ` jest niewystarczająca do tego zadania. Odmawiam wykonania!"}`)
}

func SWIRole(name string, role string) {
	slackWebhook(`{"text":"Rola ` + name + ` to: ` + role + `"}`)

}

func SWInvalidBuildParameters() {
	slackWebhook(`{"text":"Niepoprawna ilość parametrów do wykonania buildu!"}`)
}
