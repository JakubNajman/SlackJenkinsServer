package server

import (
	"log"
	"main/jenkinsbuilds"
	"main/slackapi"
	"strings"
)

func contentParser(body []string) []string {
	textContent := strings.Split(body[8], "=")[1]
	textContentArray := strings.Split(textContent, "+")
	return textContentArray
}

func BuildParser(bodyParsed []string) jenkinsbuilds.BuildJob {
	if len(bodyParsed) != 4 {
		log.Println("Invalid amout of parameters: ")
		log.Println(len(bodyParsed))
		slackapi.SWInvalidBuildParameters()
		return jenkinsbuilds.BuildJob{}
	} else {
		buildJobBody := jenkinsbuilds.BuildJob{
			Module:     bodyParsed[0],
			Branch:     bodyParsed[1],
			Version:    bodyParsed[2],
			Repository: bodyParsed[3],
		}
		return buildJobBody
	}
}

func DeployParser(bodyParsed []string) jenkinsbuilds.DeployJob {
	if len(bodyParsed) != 4 {
		log.Println("Invalid amout of parameters: ")
		log.Println(len(bodyParsed))
		slackapi.SWInvalidBuildParameters()
		return jenkinsbuilds.DeployJob{}
	} else {
		deployJobBody := jenkinsbuilds.DeployJob{
			Project:     bodyParsed[0],
			Module:      bodyParsed[1],
			Version:     bodyParsed[2],
			Environment: bodyParsed[3],
		}
		return deployJobBody
	}
}
