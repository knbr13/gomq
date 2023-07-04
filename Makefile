run:
	go run main.go

postgres:
	docker run --name fiber-container -p 5432:5432 -e POSTGRES_PASSWORD=fiber777! -d postgres:12-alpine
