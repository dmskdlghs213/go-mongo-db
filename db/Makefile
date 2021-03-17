build:
	docker build -t mongo-db .

run: build
	docker run --rm -p 27000:27017 \
	--name mongo-db \
	mongo-db

