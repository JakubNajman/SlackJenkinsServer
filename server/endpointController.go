package server

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"main/database"
	"main/jenkinsbuilds"
	"main/slackapi"
	"net/http"
	"strconv"
	"time"
)

func CheckHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.URL.Path != "/check" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Status: OK!")
}

func BuildHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	isAuth, body := slackapi.SignatureValidatorAndBodyReturner(w, r, ps)
	if !isAuth {
		log.Println("Invalid signature! Job dismissed.")
	} else {
		log.Println("Signature is correct! Job in progress.")

		content := slackapi.TextParserSlackResponse(body)
		contentParsed := contentParser(content)
		buildParsed := BuildParser(contentParsed)

		httpRequestBuild(buildParsed)
	}

}

func BuildnaskHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	isAuth, body := slackapi.SignatureValidatorAndBodyReturner(w, r, ps)
	if !isAuth {
		log.Println("Invalid signature! Job dismissed.")
	} else {
		UserId, UserName := getUserInfoFromRequest(body)
		if compareUserRole(UserId, 3) {
			log.Println("Signature is correct! Job in progress.")
			var moduleApi string = "api"
			var moduleWeb string = "web"
			var moduleWrk string = "wrk"

			var branch string = "release-nask"
			var version string = "release-nask-" + strconv.FormatInt(time.Now().Unix(), 10)
			var gcr string = "nask-stg"
			_, UserName := getUserInfoFromRequest(body)

			buildJobBodyApi := jenkinsbuilds.BuildJob{
				Module:     moduleApi,
				Branch:     branch,
				Version:    version,
				Repository: gcr,
			}

			buildJobBodyWeb := jenkinsbuilds.BuildJob{
				Module:     moduleWeb,
				Branch:     branch,
				Version:    version,
				Repository: gcr,
			}

			buildJobBodyWrk := jenkinsbuilds.BuildJob{
				Module:     moduleWrk,
				Branch:     branch,
				Version:    version,
				Repository: gcr,
			}
			httpRequestBuild(buildJobBodyApi)
			slackapi.SWBuildNaskStarted(UserName, buildJobBodyApi.GetModule(), buildJobBodyApi.GetVersion(), buildJobBodyApi.GetRepository())
			httpRequestBuild(buildJobBodyWrk)
			slackapi.SWBuildNaskStarted(UserName, buildJobBodyWrk.GetModule(), buildJobBodyWrk.GetVersion(), buildJobBodyWrk.GetRepository())
			httpRequestBuild(buildJobBodyWeb)
			slackapi.SWBuildNaskStarted(UserName, buildJobBodyWeb.GetModule(), buildJobBodyWeb.GetVersion(), buildJobBodyWeb.GetRepository())
		} else {
			slackapi.SWInsufficientRole(UserName)
		}
	}
}

func DeployHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	isAuth, body := slackapi.SignatureValidatorAndBodyReturner(w, r, ps)
	if !isAuth {
		log.Println("Invalid signature! Job dismissed.")
	} else {
		log.Println("Signature is correct! Job in progress.")

		content := slackapi.TextParserSlackResponse(body)
		contentParsed := contentParser(content)
		deployParsed := DeployParser(contentParsed)

		httpRequestDeploy(deployParsed)
	}
}

func DeploynaskHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	isAuth, body := slackapi.SignatureValidatorAndBodyReturner(w, r, ps)
	if !isAuth {
		log.Println("Invalid signature! Job dismissed.")
	} else {
		UserId, UserName := getUserInfoFromRequest(body)
		if compareUserRole(UserId, 2) {
			log.Println("Nask deployment in progress.")
			client := database.Initialisation()
			m := database.GetBuild(client, "api")
			var version string = m["Version"].(string)
			var environment string = "stg"
			var project string = "nask"
			_, UserName := getUserInfoFromRequest(body)

			deployJobBodyApi := jenkinsbuilds.DeployJob{
				Project:     project,
				Module:      "api",
				Version:     version,
				Environment: environment,
			}

			deployJobBodyWrk := jenkinsbuilds.DeployJob{
				Project:     project,
				Module:      "wrk",
				Version:     version,
				Environment: environment,
			}

			deployJobBodyWeb := jenkinsbuilds.DeployJob{
				Project:     project,
				Module:      "web",
				Version:     version,
				Environment: environment,
			}
			defer client.Close()
			httpRequestDeploy(deployJobBodyApi)
			slackapi.SWDeploy(UserName, deployJobBodyApi.GetModule(), deployJobBodyApi.GetVersion(), deployJobBodyApi.GetEnvironment())
			httpRequestDeploy(deployJobBodyWrk)
			slackapi.SWDeploy(UserName, deployJobBodyWrk.GetModule(), deployJobBodyWrk.GetVersion(), deployJobBodyWrk.GetEnvironment())
			httpRequestDeploy(deployJobBodyWeb)
			slackapi.SWDeploy(UserName, deployJobBodyWeb.GetModule(), deployJobBodyWeb.GetVersion(), deployJobBodyWeb.GetEnvironment())
		} else {
			slackapi.SWInsufficientRole(UserName)
		}

	}

}

func BuildendedHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var module string = ps.ByName("module")
	client := database.Initialisation()
	m := database.GetBuild(client, module)
	var version string = m["Version"].(string)
	var repository string = m["Repository"].(string)
	var branch string = m["Branch"].(string)
	defer client.Close()
	slackapi.SWBuildNaskEnded(module, version, repository, branch)
}

func TestHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	body := slackapi.BodyParser(w, r)
	UserId, _ := getUserInfoFromRequest(body)
	//if compareUserRole(UserId, 1) {
	log.Println(body)
	log.Println(UserId)
	//} else {
	//	slackapi.SWInsufficientRole(UserName)
	//}

}

func MyRoleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	isAuth, body := slackapi.SignatureValidatorAndBodyReturner(w, r, ps)
	if !isAuth {
		log.Println("Invalid signature! Job dismissed.")
	} else {
		client := database.Initialisation()
		UserID, UserName := getUserInfoFromRequest(body)
		role := database.GetUserRole(client, UserID)
		roleString := database.GetUserRoleFromDocument(client, role)
		defer client.Close()
		slackapi.SWIRole(UserName, roleString)
	}
}
