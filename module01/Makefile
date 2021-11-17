build:
	docker build -t httpserver .

run: build
	docker run -d -p 8080:8080 httpserver