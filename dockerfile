FROM golang:1.16

WORKDIR /usr/test
COPY go.mod .

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go get -u github.com/jackc/tern
RUN go get -u github.com/gin-gonic/gin

COPY . .
RUN go build main.go

EXPOSE 8080


RUN ls -latr
#RUN chmod 755 entrypoint.sh
# CMD ls -latr
#ENTRYPOINT ["./entrypoint.sh"]
CMD go run main.go
