FROM golang:latest

RUN mkdir Gotties-Server

COPY src /go/src/Gotties-Server/src

WORKDIR /go/src/Gotties-Server/src


RUN apt-get update \
  && apt-get install -y mariadb-client \
#  && go get -u github.com/gin-gonic/gin \
#  && go get github.com/jinzhu/gorm \
#  && go get github.com/go-sql-driver/mysql \
#  && go get aws.google.com/go \
#  && go get github.com/joho/godotenv \
  && GOOS=linux GOARCH=amd64 go build main.go

CMD ["sh", "./start_app.sh"]
#CMD ["./main"]
