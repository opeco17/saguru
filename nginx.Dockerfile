FROM  nginx:latest

ADD nginx/conf/default.conf /etc/nginx/conf.d/default.conf

COPY nginx/conf/certs/fullchain.pem /etc/nginx/certs/fullchain.pem
COPY nginx/conf/certs/privkey.pem /etc/nginx/certs/privkey.pem
