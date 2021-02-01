FROM golang:1.15 as builder

WORKDIR /src/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN GOOS=linux CGO_ENABLED=0 go build main.go

EXPOSE 8080

CMD ["/src/main"]
