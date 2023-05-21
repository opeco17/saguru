FROM golang:1.18

RUN go install github.com/cosmtrek/air@latest

WORKDIR /usr/src/app/lib/
COPY lib/go.mod lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY api/go.mod api/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY . ./app

WORKDIR /usr/src/app/api

CMD ["air"]