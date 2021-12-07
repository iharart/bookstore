##
## Build
##
FROM golang:1.16-buster AS build
#FROM golang:1.16-alpine AS build

WORKDIR ./src/bookstore
COPY . .

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./app/ ./
COPY ./cmd/ ./

RUN go mod download

RUN go get -u github.com/joho/godotenv
RUN go get -u gorm.io/gorm

RUN go get -u github.com/go-sql-driver/mysql
RUN go get gorm.io/driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/iharart/bookstore
RUN go get -u github.com/sirupsen/logrus
RUN go get -u github.com/stretchr/testify/assert
RUN go get -u github.com/stretchr/testify/suite
RUN go get -u github.com/ory/dockertest/v3
RUN go get -u github.com/golangci/golangci-lint/

RUN go build -o /bookstore

##
## Deploy
##

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /bookstore /bookstore
EXPOSE 8082

CMD ["/bookstore"]