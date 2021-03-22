# grpc

Package grpc is a very simple introduction to creating a GRPC API definition using a `.proto` file and generating code for it using the `protoc` compiler.

## Setting up protoc and Go code generators

The `protoc` compiler can be downloaded from [https://github.com/protocolbuffers/protobuf/releases/](https://github.com/protocolbuffers/protobuf/releases/).

We need to install two Go plugins, which generate Go structs for protocol buffer messages and scaffolding for a GRPC server and client respectively. This can be done by running the following command:

```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go \
       google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

This command will install the relevant executables to either $GOBIN or ${GOPATH}/bin. The directory should be on your PATH for the `protoc` compiler to find them.

## Generating core from the .proto file

The following command will generate both the structs and the scaffolding.

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       contacts.proto
```

This will generate Go source files in the same directory. The file *contacts.pb.go* will contain structs derived from the messages defined in the .proto file. The file *contacts_grpc.pb.go* will contain a fully implemented GRPC client for the service defined in the .proto file. It will also contain scaffolding for implementing a GRPC server for the service, notably an interface that has to be implemented when writing the actual GRPC server.
