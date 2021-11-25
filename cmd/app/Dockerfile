FROM golang:1.16-alpine as build-env

WORKDIR ./bookstore

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY *.go ./

RUN ls
RUN pwd

RUN go mod download

RUN go get -u github.com/joho/godotenv
RUN go get -u gorm.io/gorm

RUN go get -u github.com/go-sql-driver/mysql
RUN go get gorm.io/driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/iharart/app
RUN go get -u github.com/sirupsen/logrus
RUN go get -u github.com/stretchr/testify/assert
RUN go get -u github.com/stretchr/testify/suite
RUN go get -u github.com/ory/dockertest/v3
RUN go get -u github.com/DATA-DOG/go-sqlmock
RUN apk update && apk add bash
RUN go build -o ./app

#RUN go get github.com/githubnemo/CompileDaemon

#ENTRYPOINT CompileDaemon --build "go build main.go" --comand=./main
#COPY ../.env ../.

EXPOSE 8082

RUN ls
CMD ["./bookstore"]
