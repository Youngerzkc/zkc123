myredis_dir=/home/younger/project/goproject/src/github.com/zkc123/database/myredis
mysql_dir=/home/younger/project/goproject/src/github.com/zkc123/database/mysql


run: run-docker-mysql run-docker-redis run-main
run-main:
	go run main.go
	#注意mysql数据的持久化
run-docker-mysql:
	docker run  -d  -p 3307:3306 -e MYSQL_ROOT_PASSWORD=123456 -v $(mysql_dir):/var/lib/mysql -e MYSQL_USERNAME=root mysql:5.7.4 
	#注意数据的持久化
run-docker-redis:	
	docker run  -d  -v $(myredis_dir):/data  -p 6381:6379 redis
build-images:
	
run-docker:

