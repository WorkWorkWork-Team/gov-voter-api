run:
	docker run -i -t -p=3306:3306 dev/href-counter:latest
#	ENV=dev go run main.go

start-dev-db-linux:
	docker run --name mysql-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v ${PWD}/database:/docker-entrypoint-initdb.d \
        mysql:latest

#not try yet
start-dev-db-window:
	docker run --name mysql-dev -d \
        -p 3306:3306 \
        -e MYSQL_ROOT_PASSWORD=P@ssw0rd \
        --restart unless-stopped \
        -v %cd%/database:/docker-entrypoint-initdb.d \
        mysql:latest

build-php-admin:
	docker run --name phpmyadmin -d --link mysql-dev:db -p 8080:80 phpmyadmin/phpmyadmin