# TOS
This is a distributed ticket order system. There are various commands located in `Makefile` for building and starting the server, as well as in `scripts` are helper scripts.

The server utilizes gRPC allowing clients to make calls, i.e. requesting menu items, submitting an order. The idea is for a front client to implement the `MenuService` fully, and at least the `OrderService.PublishOrder` function. A back client then implements the rest of `OrderService` and subscribes for orders. Both clients are still a work in progress while the server is stable.

### Building and running
```
make
make start
```

When running the server, supply a database path or it defaults to tmp storage.

### Debugging
`export GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info GODEBUG=http2debug=2`
