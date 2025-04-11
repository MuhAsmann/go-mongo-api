run:
	sudo docker compose up -d
build:
	sudo docker compose build
stop:
	sudo docker compose down
restart:
	sudo docker compose down && sudo docker compose build && sudo docker compose up -d
logs:
	sudo docker compose logs -f
shell: