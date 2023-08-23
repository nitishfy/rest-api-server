package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/nitishfy/rest-api-server/swagger/server/restapi"
	"github.com/nitishfy/rest-api-server/swagger/server/restapi/operations"
	"log"
	"net/http"
)

func main() {
	// Initialize swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}
	api := operations.NewHelloAPIAPI(swaggerSpec)
	server := restapi.NewServer(api)

	defer func() {
		if err := server.Shutdown(); err != nil {
			log.Fatalln(err)
		}
	}()
	server.Port = 8081
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(GetHandler)
	api.GetAPIAPIVersionHandler = operations.GetAPIAPIVersionHandlerFunc(GetApiVersion)
	api.GetUsersUserHandler = operations.GetUsersUserHandlerFunc(GetUsers)
	api.GetImagesNameHandler = operations.GetImagesNameHandlerFunc(GetImage)

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

func GetHandler(operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("Hello World")
}

func GetUsers(params operations.GetUsersUserParams) middleware.Responder {
	return operations.NewGetUsersUserOK().WithPayload(params.User)
}

func GetApiVersion(params operations.GetAPIAPIVersionParams) middleware.Responder {
	return operations.NewGetAPIAPIVersionOK().WithPayload(params.APIVersion)
}

func GetImage(params operations.GetImagesNameParams) middleware.Responder {
	var URL string

	if params.Name != "" {
		URL = "https://github.com/scraly/gophers/raw/main/" + params.Name + ".png"
	} else {
		URL = "https://github.com/scraly/gophers/raw/main/dr-who.png"
	}

	response, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}
	return operations.NewGetImagesNameOK().WithPayload(response.Body)
}
