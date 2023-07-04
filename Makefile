CONTAINER_NAME=fiber-container

run:
	go run main.go

postgres:
	docker run --name ${CONTAINER_NAME} -p 5432:5432 -e POSTGRES_PASSWORD=fiber777! -d postgres:12-alpine

createdb:
	docker exec -it ${CONTAINER_NAME} createdb -U postgres gorm
