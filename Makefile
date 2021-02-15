init:
	go get ./...
	
# comandos para execução

run:
	go run main.go

run-watch:
	nodemon --exec go run main.go --signal SIGTERM

# comandos para documentação
docs:
	swag init