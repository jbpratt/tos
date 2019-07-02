FROM golang:1.12

RUN mkdir /src

ADD . /src/

WORKDIR /src

RUN go get -d -v ./...

RUN go build -o main .

EXPOSE 9001
EXPOSE 50051

ENTRYPOINT ["./main"]
