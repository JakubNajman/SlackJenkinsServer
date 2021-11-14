package server

import (
	f "fmt"
	"io"
	"main/jenkinsbuilds"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func requestExecutor(data url.Values, url string) {
	var res *http.Response = nil

	auth := jenkinsbuilds.Authentication{
		User:  os.Getenv("LOGIN"),
		Token: os.Getenv("TOKEN"),
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(data.Encode()))

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	f.Println("URL:>", url)
	f.Println("BODY:>", data)

	req.SetBasicAuth(auth.User, auth.Token)

	client := &http.Client{}

	res, e := client.Do(req)

	if e != nil {
		panic(e)
	}
	f.Println("RESPONSE STATUS:>", res.Status)

	f.Println("USER:>", auth.GetUser())
	f.Println("TOKEN:>", auth.GetToken())

	io.Copy(os.Stdout, res.Body)
}
