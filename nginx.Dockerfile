FROM  nginx:latest

ADD nginx/conf/default.conf /etc/nginx/conf.d/default.conf

COPY nginx/conf/certs/ /etc/nginx/certs/