FROM golang:ubuntu20.04

LABEL MAINTAINER="Ari Satrio<ari.satrio@uii.ac.id>"

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Install library auto-reload
RUN go install github.com/cosmtrek/air@latest

COPY . ./

RUN go build -o /main

CMD [ "air" ]
