FROM mysql


ENV MYSQL_ROOT_PASSWORD=secret \MYSQL_DATABASE=Academia 

ADD academia.sql /docker-entrypoint-initdb.d/

EXPOSE 3306:3306