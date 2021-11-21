FROM --platform=linux/x86_64 mysql:8.0

ENV MYSQL_ROOT_PASSWORD=root
ENV TZ=Asia/Tokyo
ENV BIND-ADDRESS=0.0.0.0

COPY ./database/mysql_init /docker-entrypoint-initdb.d
COPY ./database/conf.d /etc/mysql/conf.d