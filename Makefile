
run-go:
	go run command/main.go

run-psql:
	sudo docker start mediumuz

run-redis:
	sudo docker start redisdb

start-psql:
	sudo docker run --name mediumuz -e POSTGRES_PASSWORD=0224 -d -p 2001:5432 postgres

start-redis:
	sudo docker run --redisdb redis-test-instance -p 6379:6379 -d redis
