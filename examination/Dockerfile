FROM golang:latest

WORKDIR $GOPATH/src/examination
COPY . $GOPATH/src/examination
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./examination"]