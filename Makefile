run:
	ENV=dev go run main.go

start-dev-db-linux:
	docker start mysql-dev || docker run --name mysql-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v ${PWD}/database:/docker-entrypoint-initdb.d \
        mysql:latest

#not try yet
start-dev-db-window:
	docker start mysql-dev || docker run --name mysql-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v %cd%/database:/docker-entrypoint-initdb.d \
        mysql:latest

start-dev-php-admin:
	docker start phpmyadmin || docker run --name phpmyadmin -d --link mysql-dev:db -p 8080:80 phpmyadmin/phpmyadmin

unit-test:
	ginkgo -r
