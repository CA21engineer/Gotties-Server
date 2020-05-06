FROM golang:latest

RUN mkdir Gotties-Server

COPY src /go/src/Gotties-Server/src

WORKDIR /go/src/Gotties-Server/src


RUN apt-get update \
  && apt-get install -y mariadb-client \
  && GOOS=linux GOARCH=amd64 go build main.go

CMD ["sh", "./start_app.sh"]
#CMD ["./main"]
