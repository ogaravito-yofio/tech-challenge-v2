package main

import (
	"log"
	"os"
	"strconv"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/ogaravito-yofio/tech-challenge-v2/gitlab.com/yofio_v2/tech_challenge_v2"
	"github.com/ogaravito-yofio/tech-challenge-v2/gitlab.com/yofio_v2/tech_challenge_v2/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(tech_challenge_v2.SwaggerJSON, tech_challenge_v2.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewTechChallengeV2API(swaggerSpec)
	server := tech_challenge_v2.NewServer(api)
	defer func() {
		_ = server.Shutdown()
	}()

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "YoFio Recruitment - Tech Challenge"
	parser.LongDescription = swaggerSpec.Spec().Info.Description
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	port := os.Getenv("PORT")
	log.Printf("PORT loaded from env var: [%s]", port)
	if port != "" {
		server.Host = ""
		server.Port, err = strconv.Atoi(port)
		if err != nil {
			log.Fatalln(err)
		}
	}

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
