package server

import (
	"main/database"
	"main/jenkinsbuilds"
	"net/url"
	"os"
)

func httpRequestBuild(body jenkinsbuilds.BuildJob) {
	data := url.Values{}

	data.Set("VERSION", body.GetVersion())
	data.Set("BRANCH", body.GetBranch())
	data.Set("REPOSITORY", body.GetRepository())

	var jenkinsUrl string = os.Getenv("JENKINS_URL")

	var url string = "http://" + jenkinsUrl + ":9999/job/Build/job/" + body.GetModule() + "/buildWithParameters"

	requestExecutor(data, url)
	client := database.Initialisation()
	database.UpdateBuild(client, body)
}

func httpRequestDeploy(body jenkinsbuilds.DeployJob) {
	data := url.Values{}

	data.Set("VERSION", body.GetVersion())
	data.Set("ENVIRONMENT", body.GetEnvironment())

	var url string = "http://10.8.0.1:9999/job/Deploy/job/" + body.GetProject() + "/job/" + body.GetModule() + "/buildWithParameters"
	requestExecutor(data, url)
}
