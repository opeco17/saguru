FROM --platform=linux/x86_64 golang:1.16.5

WORKDIR /usr/src/app/lib/
COPY backend/lib/go.mod backend/lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY backend/job/go.mod backend/job/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY backend ./app

WORKDIR /usr/src/app/job