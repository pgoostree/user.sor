# User System Of Record (SOR) API

## Description
This is a User and UserGroup management REST API that provides the following capabilities:

## Users:
- Add a user 
- Retrieve a user 
- Update a user 
- Delete a user 

## Groups:
- Add a group 
- Get a list of users associated to a group
- Update the users associated with the group
- Delete the group

## Pre-requisites
- Make for running the Makefile commands
- Docker to run the API and Postgres database together  [Download link](https://docs.docker.com/get-docker/)

## Run unit tests
```bash
make unit-tests
```
## Run integration tests
```bash
make integration-tests
```

## Run the service
```bash
make run
```

### Create some users

```bash
curl --location --request POST 'http://localhost:9000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_id" : "theDude",
	"first_name" : "Jeff",
	"last_name" : "Lebowski"
}'

curl --location --request POST 'http://localhost:9000/users' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_id" : "donnie",
	"first_name" : "Theodore",
	"last_name" : "Kerabatsos"
}'
```

### Create a group
```bash
curl --location --request POST 'http://localhost:9000/groups' \
--header 'Content-Type: application/json' \
--data-raw '{
	"name" : "TheDudesTeam"
}'
```

### Add users to a group
```bash
curl --location --request PUT 'http://localhost:9000/groups/TheDudesTeam' \
--header 'Content-Type: application/json' \
--data-raw '{
	"user_ids" : [
        "theDude", 
        "donnie"
    ]
}'
```

### Get the users in a group
```bash
curl --location --request GET 'http://localhost:9000/groups/TheDudesTeam'
```

## Stop the service
```bash
make stop
```

### Run the service locally in development environment

This service API is written in Go and the database is Postgres.

By default when running locally in the development environment it is configured to connect to Postgres at localhost:5432.

To run the service locally in dev you need to first start the database:

```bash
docker-compose up postgres
```

And then start the service with the following command in the root directory of the project:

```base
go run main.go
```
