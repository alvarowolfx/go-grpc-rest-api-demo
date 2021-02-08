# Rest API using gRPC - Talk

We are gonna create an API simulating a Movie and TV Show catalog that people can create an account and mark their favorite things on the catalog.

### Running locally

To run on dev with hot reload, install [air](https://github.com/cosmtrek/air) and run:

```
air -c air.toml
```

The project also uses a Badger as the local database.

### Generating code from our .proto file

We are gonna use [grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway) to also expose our gRPC service as a Rest API to be used by the Frontend apps. To install the extra deps, use this commmand.

```
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

To generate the files from our .proto run:

```
make protoc-generate
```

### gRPC

Our service will expose a gRPC API.

### Rest API and Swagger

The same .proto file has annotations to generate a gateway/bridge between Rest and gRPC. Those annotation are provided by the `grpc-gateway` project mentioned before.
