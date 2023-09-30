start-dev:
	go run .


docker-build:
	docker build -t go-web-server-template .

docker-start:
	docker run -p 10000:10000 go-web-server-template
