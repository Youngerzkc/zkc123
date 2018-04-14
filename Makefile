run:
	go run main.go
	#注意mysql数据的持久化
	docker run  -d  -p 3307:3306  -e MYSQL_ROOT_PASSWORD=123456 mysql:5.7.4
	#注意数据的持久化
	docker run  -d  -p 6380:6379 redis 
build-images:

run-docker:
