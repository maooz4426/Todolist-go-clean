setup:
	go get github.com/stretchr/testify

up:
	docker compose up

build:
	docker compose build

build-up:
	docker compose up --build

down:
	docker compose down

mysql:
	mysql -h 127.0.0.1 -u user -ppassword db

gomock:
	mockgen -source=domain/repository/todo.go -destination=domain/mockreposiotry/todo.go -package mockrepository & mockgen -source=domain/repository/transaction_manager.go -destination=domain/mockreposiotry/transaction_manager.go -package mockrepository

