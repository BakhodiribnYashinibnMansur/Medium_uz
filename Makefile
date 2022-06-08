run-all:
	sudo docker start mediumuz && sudo docker start redisdb && go run command/main.go

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

swag:
	swag init -g  command/main.go

migrate-up:
	migrate -path ./schema -database 'postgres://quffklnl:YuXp2n2HtszNTlB85lE75QJ8xZ0aLbXP@heffalump.db.elephantsql.com/quffklnl?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://quffklnl:YuXp2n2HtszNTlB85lE75QJ8xZ0aLbXP@heffalump.db.elephantsql.com/quffklnl?sslmode=disable' down

docker-image:
	sudo docker build -t mediumuz .

docker-container:
	sudo docker run -it --name heroku2  -p 8080:8080 medium

deploy-heroku:
	git add . && git commit -m "set-up heroku" && git push origin main && git push heroku main