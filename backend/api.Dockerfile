FROM golang:1.16.5

RUN go get github.com/cosmtrek/air

WORKDIR /usr/src/app/lib/
COPY lib/go.mod lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY api/go.mod api/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY . ./app

WORKDIR /usr/src/app/api

CMD ["air"]