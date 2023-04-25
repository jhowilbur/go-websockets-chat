#FROM golang:1.19-buster as builder
#
#WORKDIR /app
#
#COPY . ./
#RUN go mod download
#
#RUN go build -v cmd/web/*.go -o server

FROM debian:buster-slim

COPY ./html ./html
COPY ./main ./main

EXPOSE 5000

# Run the web service on container startup.
CMD ["/main"]
