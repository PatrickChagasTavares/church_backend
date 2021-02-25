.PHONY: run watch docker-up docker-down mocks test coover docs

# Styles
red=\033[31m
green=\033[32m
yellow=\033[33m
blue=\033[34m
purple=\033[35m
cyan=\033[36m
blink=\033[5m
reset=\033[0m

# Variables

su :=

# comandos para subir o docker-compose

docker-up: 
	@echo "${green}\n\nInicializando instancias${reset}\n\n"
	docker-compose -- up -d

docker-down:
	@echo "${green}\n\nDesmontando instancias${reset}\n\n"
	$(su)docker-compose down --remove-orphans -v

# comandos para execução
run:
	go run main.go

run-watch:
	nodemon --exec go run main.go --signal SIGTERM

# comandos para documentação
docs:
	swag init

# comandos para teste
test:
	go test -v -p 1 -cover -failfast ${FOLDER_TEST}  -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

test-cover: test
	go tool cover -html=coverage.out

mocks: 
	rm -rf ./mocks
	mkdir mocks

	mockgen -source=./store/health/health.go -destination=./mocks/health_mock.go -package=mocks -mock_names=Store=MockHealthStore
	mockgen -source=./store/children/children.go -destination=./mocks/children_mock.go -package=mocks -mock_names=Store=MockChildrenStore