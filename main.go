package main

import (
	"fmt"
	"log"
	"main/server"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("Starting Golang Middleman Server (GMS) at port 9191.")

	router := httprouter.New()
	router.GET("/check", server.CheckHandler)
	router.GET("/build/:module/:branch/:version/:gcr", server.BuildHandler)
	router.GET("/deploy/:project/:module/:version/:environment", server.DeployHandler)
	router.POST("/buildnask", server.BuildnaskHandler)
	router.POST("/deploynask", server.DeploynaskHandler)
	router.POST("/buildended/:module", server.BuildendedHandler)
	router.POST("/test", server.TestHandler)
	router.POST("/myrole", server.MyRoleHandler)

	log.Fatal(http.ListenAndServe(":9191", router))

}
