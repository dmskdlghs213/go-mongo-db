up: build-go
	docker-compose up go-server
build: 
	docker-compose build go-server
	docker-compose build mongo-server
build-go:
	docker-compose build go-server
build-mongo:
	docker-compose build mongo-server