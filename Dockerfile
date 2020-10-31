FROM golang as builder
ENV GO111MODULE=on
WORKDIR /code
ADD . .
RUN make server

FROM golang:latest
EXPOSE 9001
EXPOSE 50051
WORKDIR /
COPY --from=builder /code/bin/server /usr/bin/server
ENTRYPOINT ["/usr/bin/server"]
