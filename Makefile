unit-test:
	go test -v ./...      

build:
	go build main.go

build-docker:
	docker build -t user.sor .

run:
	docker-compose up -d

stop:
	docker-compose down
