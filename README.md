# Cron

Project to generate cron executions for YoFio projects.

## Generate code for the server from swagger

`$` `go mod init github.com/ogaravito-yofio/tech-challenge-v2`
`$` `swagger generate server -s gitlab.com/yofio_v2/tech-challenge-v2 -f swagger.yaml -A tech-challenge-v2`

## Run

For local environment just execute `make run-local`
