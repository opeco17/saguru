FROM --platform=linux/x86_64 golang:1.16.5

RUN go get github.com/cosmtrek/air

WORKDIR /usr/src/app/lib/
COPY backend/lib/go.mod backend/lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY backend/api/go.mod backend/api/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY backend ./app

WORKDIR /usr/src/app/api

CMD ["air"]