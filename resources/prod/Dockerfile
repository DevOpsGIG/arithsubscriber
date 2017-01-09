FROM golang:1.8-alpine

COPY . src/github.com/devopsgig/arithsubscriber

WORKDIR src/github.com/devopsgig/arithsubscriber

RUN go build

CMD ["./arithsubscriber"]

EXPOSE 8081
