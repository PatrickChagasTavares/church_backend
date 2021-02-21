.PHONY: run run-watch docker-up docker-down docs

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