build:
	docker-compose build

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

rmi:
	docker rmi $(docker images -a -q)

rm:
	docker rm $(docker ps -a -f status=exited -q)
