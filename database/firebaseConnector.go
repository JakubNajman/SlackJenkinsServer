package database

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	_ "fmt"
	"github.com/fatih/structs"
	"google.golang.org/api/option"
	"log"
	"main/jenkinsbuilds"
	"strconv"
)

func Initialisation() *firestore.Client {
	opt := option.WithCredentialsFile("/go/src/gms/database/gms-jenkins-firebase-key.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}

func UpdateBuild(client *firestore.Client, job jenkinsbuilds.BuildJob) {
	log.Println("module ", job.GetModule())
	log.Println(structs.Map(job))
	result, err := client.Collection("builds").Doc(job.GetModule()).Set(context.Background(), structs.Map(job))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)

}

func GetBuild(client *firestore.Client, module string) map[string]interface{} {
	result, err := client.Collection("builds").Doc(module).Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	m := result.Data()
	return m
}

func GetUserRole(client *firestore.Client, UserID string) int {
	result, err := client.Collection("users").Doc("users").Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	m := result.Data()

	userRoleString, ok := m[UserID].(string)
	if ok {
		log.Println(userRoleString)

		res, _ := strconv.Atoi(userRoleString)

		return res
	} else {
		return 100
	}
}

func GetUserRoleFromDocument(client *firestore.Client, role int) string {
	result, err := client.Collection("users").Doc("roles").Get(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	m := result.Data()
	userRoleString, ok := m[strconv.Itoa(role)].(string)
	if ok {
		log.Println(userRoleString)

		return userRoleString
	} else {
		return ""
	}

}
