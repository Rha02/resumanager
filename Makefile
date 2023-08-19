SHELL=cmd.exe
BACKEND_BINARY=backend

up:
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started

up_build: build_backend
	@echo Stopping docker images... (if running)
	docker-compose down
	@echo Building (if required) and starting docker images
	docker-compose up -d --build
	@echo Docker images started

down:
	@echo Stopping docker images...
	docker-compose down
	@echo Docker images stopped

build_backend:
	@echo Building backend binary...
	chdir backend && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${BACKEND_BINARY} ./cmd/web
	@echo Backend binary built
