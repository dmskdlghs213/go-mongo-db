build:
	docker build -t go-mongo-db .

run: build
	docker run --rm -p 9000:8888 \
	--env-file .env \
	--name go-mongo-db \
	go-mongo-db