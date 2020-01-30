FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code
ADD go.mod go.sum /code/
RUN go mod download
ADD . .
RUN go build -o /app server.go 

FROM golang:1.12 
EXPOSE 9001
EXPOSE 50051
WORKDIR /
COPY --from=builder /app /usr/bin/app
ENTRYPOINT ["/usr/bin/app"]
