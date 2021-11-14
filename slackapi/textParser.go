package slackapi

import (
	"strings"
)

func TextParserSlackResponse(body string) []string {
	bodySplitted := strings.Split(body, "&")
	return bodySplitted
}
