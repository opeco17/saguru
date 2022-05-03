FROM golang:1.16.5 AS api-build

WORKDIR /usr/src/app/lib/
COPY lib/go.mod lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY api/go.mod api/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY . ./app

WORKDIR /usr/src/app/api
RUN go build -o /usr/local/go/bin/api


FROM gcr.io/distroless/base
COPY --from=api-build /usr/local/go/bin/api /
ENTRYPOINT ["/api"]