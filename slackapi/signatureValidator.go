package slackapi

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func Abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func timeValidator(timestamp int64) bool {
	if Abs(time.Now().Unix()-timestamp) > 60*5 {
		return false
	} else {
		return true
	}

}

func BodyParser(w http.ResponseWriter, r *http.Request) string {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return ""
	}
	return string(bodyByte[:])
}

func SignatureValidatorAndBodyReturner(w http.ResponseWriter, r *http.Request, ps httprouter.Params) (bool, string) {
	var slackSecret string = os.Getenv("SLACK_SECRET")

	body := BodyParser(w, r)

	log.Println("Reading body from request:")
	log.Println(body)

	timestamp := r.Header.Get("X-Slack-Request-Timestamp")
	n, _ := strconv.ParseInt(timestamp, 10, 64)

	if !timeValidator(n) {
		return false, ""
	} else {
		var sigBasestring string = "v0:" + timestamp + ":" + body

		h := hmac.New(sha256.New, []byte(slackSecret))
		h.Write([]byte(sigBasestring))
		sha := "v0=" + hex.EncodeToString(h.Sum(nil))

		shaRequest := r.Header.Get("X-Slack-Signature")

		if sha == shaRequest {
			return true, body
		} else {
			return false, ""
		}
	}
}
