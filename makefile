up:
	docker compose up

build:
	docker compose build

build-up:
	docker compose up --build

down:
	docker compose down

mysql:
	mysql -h 127.0.0.1 -u user -p db