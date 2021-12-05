FROM  nginx:latest

ADD conf/default_prod.conf /etc/nginx/conf.d/default.conf

COPY conf/certs/ /etc/nginx/certs/