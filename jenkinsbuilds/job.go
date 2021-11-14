package jenkinsbuilds

import (
	"encoding/json"
	"log"
)

//Jenkins Job Interface

//TODO 1.1 Find way to implement interfaces without method redundancy.

// type JenkinsJobInterface interface {
// 	getDetails() string
// 	getBranch() string
// 	getModule() string
// 	getRepository() string
// 	getVersion() string
// 	getEnvironment() string
// }

//Jenkins Build Job Struct - Init & Methods

type BuildJob struct {
	Module     string `json:"module"`
	Branch     string `json:"branch"`
	Repository string `json:"gcr"`
	Version    string `json:"version"`
}

// func newBuild(bj BuildJob) JenkinsJobInterface {
// 	return newBuildJob(bj)
// }

// func newBuildJob(bj BuildJob) JenkinsJobInterface {
// 	var convertedBuildJobToJobInterface JenkinsJobInterface = bj
// 	return convertedBuildJobToJobInterface
// }

func (bj *BuildJob) SetModule(module string) {
	bj.Module = module
}

func (bj *BuildJob) SetVersion(version string) {
	bj.Version = version
}

func (bj *BuildJob) SetBranch(branch string) {
	bj.Branch = branch
}

func (bj *BuildJob) SetRepository(repository string) {
	bj.Repository = repository
}

func (bj BuildJob) GetModule() string {
	return bj.Module
}

func (bj BuildJob) GetVersion() string {
	return bj.Version
}

func (bj BuildJob) GetBranch() string {
	return bj.Branch
}

func (bj BuildJob) GetRepository() string {
	return bj.Repository
}

//Jenkins Build Job Struct - Business Functions

func (bj BuildJob) getDetails() string {
	return bj.Module + " " + bj.Version + " " + bj.Branch + " " + bj.Repository
}

func (bj BuildJob) ToJson() []byte {
	json, err := json.Marshal(bj)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return json
}

type DeployJob struct {
	Project     string `json:"project"`
	Module      string `json:"module"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}

// func newDeploy(dj DeployJob) JenkinsJobInterface {
// 	return newDeployJob(dj)
// }

// func newDeployJob(dj DeployJob) JenkinsJobInterface {
// 	var convertedDeployJobToJobInterface JenkinsJobInterface = dj
// 	return convertedDeployJobToJobInterface
// }

func (dj *DeployJob) SetProject(project string) {
	dj.Project = project
}

func (dj *DeployJob) SetModule(module string) {
	dj.Module = module
}

func (dj *DeployJob) SetVersion(version string) {
	dj.Version = version
}

func (dj *DeployJob) SetEnvironment(environment string) {
	dj.Environment = environment
}

func (dj DeployJob) GetProject() string {
	return dj.Project
}

func (dj DeployJob) GetModule() string {
	return dj.Module
}

func (dj DeployJob) GetVersion() string {
	return dj.Version
}

func (dj DeployJob) GetEnvironment() string {
	return dj.Environment
}
