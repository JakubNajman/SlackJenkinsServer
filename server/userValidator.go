package server

import (
	"log"
	"main/database"
	"main/slackapi"
	"strings"
)

func getUserInfoFromRequest(body string) (string, string) {
	bodySplitted := slackapi.TextParserSlackResponse(body)
	log.Println("Body Splitted: " + strings.Join(bodySplitted, " "))
	return strings.Split(bodySplitted[5], "=")[1], strings.Split(bodySplitted[6], "=")[1]
}

func compareUserRole(UserID string, roleToCompare int) bool {
	client := database.Initialisation()
	userRole := database.GetUserRole(client, UserID)

	if userRole <= roleToCompare {
		return true
	} else {
		return false
	}
}
