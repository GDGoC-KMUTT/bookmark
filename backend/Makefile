dev:
	go run main.go

swag:
	swag init

build:
	env GOOS=linux GOARCH=amd64 go build -o build/backend

mock:
	mockery --all --with-expecter --dir=./internals/repositories --output=./mocks/repositories --outpkg=mockRepositories & mockery --all --with-expecter --dir=./internals/services --output=./mocks/services --outpkg=mockServices