FROM golang:latest

RUN mkdir Gotties-Server

WORKDIR /go/src/Gotties-Server/src

COPY src /go/src/Gotties-Server/src


RUN apt-get update \
  && apt-get install -y mariadb-client \
  && go get -u github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql \
  && go get firebase.google.com/go \
  && github.com/joho/godotenv \



CMD ["go", "run", "main.go"]
