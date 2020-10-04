# TOS
This is a distributed ticket order system. There are various commands located in `Makefile` for building and starting the server, as well as in `scripts` are helper scripts.

### Building and running
```
make
make start
```

When running the server, supply a database path or it defaults to tmp storage.

### Debugging
`export GRPC_GO_LOG_VERBOSITY_LEVEL=99 GRPC_GO_LOG_SEVERITY_LEVEL=info GODEBUG=http2debug=2`
