// This file is safe to edit. Once it exists it will not be overwritten

package tech_challenge_v2

import (
	"crypto/tls"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/ogaravito-yofio/tech-challenge-v2/models"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ogaravito-yofio/tech-challenge-v2/gitlab.com/yofio_v2/tech_challenge_v2/operations"
	"github.com/ogaravito-yofio/tech-challenge-v2/gitlab.com/yofio_v2/tech_challenge_v2/operations/pet"
)

//go:generate swagger generate server --target ../../../../tech-challenge-v2 --name TechChallengeV2 --spec ../../../swagger.yaml --server-package gitlab.com/yofio_v2/tech-challenge-v2 --principal interface{}

func configureFlags(_ *operations.TechChallengeV2API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TechChallengeV2API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.PetAddPeopleHandler = pet.AddPeopleHandlerFunc(func(params pet.AddPeopleParams) middleware.Responder {
		log.Printf("Receiving a request for creating new people")
		p := &models.APIResponse{
			ID: strfmt.UUID(uuid.New().String()),
		}
		return pet.NewAddPeopleCreated().WithPayload(p)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(_ *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(_ *http.Server, _, _ string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}