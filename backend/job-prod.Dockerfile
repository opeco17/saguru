FROM --platform=linux/x86_64 golang:1.16.5 AS job-build

WORKDIR /usr/src/app/lib/
COPY lib/go.mod lib/go.sum ./

WORKDIR /usr/src/app/api/
COPY job/go.mod ./job/go.sum ./
RUN go mod download

WORKDIR /usr/src/
COPY . ./app

WORKDIR /usr/src/app/job
RUN go build -o /usr/local/go/bin/job


FROM --platform=linux/x86_64 gcr.io/distroless/base
COPY --from=job-build /usr/local/go/bin/job /