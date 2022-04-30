FROM golang:1.16.5

WORKDIR /usr/src/app/lib/
COPY lib/go.mod lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY job/go.mod ./job/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY . ./app

WORKDIR /usr/src/app/job
