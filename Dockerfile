##
## Build
##

FROM golang:1.16-buster AS build

WORKDIR /app


COPY . .



RUN go build -a -o Bookstore cmd/main.go

##
## Deploy
##

FROM debian:bullseye

WORKDIR /

COPY --from=build app/Bookstore .

EXPOSE 8080

ENTRYPOINT ["/Bookstore"]