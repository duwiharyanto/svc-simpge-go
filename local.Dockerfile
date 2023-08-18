FROM arisatrio03/golang:1.17-ubuntu

LABEL MAINTAINER="Ari Satrio<ari.satrio@uii.ac.id>"

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Install library auto-reload
RUN go install github.com/cosmtrek/air@v1.42.0

COPY . ./

# RUN go build -o /main

CMD [ "air" ]
