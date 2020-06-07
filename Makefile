unit-test:

build:
	go build main.go

build-docker:
	docker build -t user.sor .

run:
	docker-compose up
