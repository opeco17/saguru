FROM --platform=linux/x86_64 golang:1.16.5

RUN go get github.com/cosmtrek/air

WORKDIR /usr/src/app/lib/
COPY app/lib/go.mod app/lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY app/api/go.mod app/api/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY app ./app

WORKDIR /usr/src/app/api

CMD ["air"]