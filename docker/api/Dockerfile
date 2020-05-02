FROM golang:latest

COPY src/api /go/src/api

WORKDIR /go/src/api/

RUN apt-get update \
  && apt-get install -y mysql-client \
  && go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

CMD ["go", "run", "main.go"]
