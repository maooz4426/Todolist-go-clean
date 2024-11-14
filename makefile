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

mysql_test:
	mysql -h 127.0.0.1 -u test_user -ptest_password test_db

gomock:
	mockgen -source=domain/repository/todo.go -destination=domain/mockreposiotry/todo.go -package mockrepository & mockgen -source=domain/repository/transaction_manager.go -destination=domain/mockreposiotry/transaction_manager.go -package mockrepository

test:
	docker exec todo_server go test -v ./...

#make gen/migration/file FILE_NAME={file_name}
gen/migration/file:
	$(eval DATE:=$(shell TZ=JST-9 date "+%Y%m%d%H%M%S"))
	touch migrations/$(DATE)-$(FILE_NAME).sql

sql/migrate:
	sql-migrate up -env="development"