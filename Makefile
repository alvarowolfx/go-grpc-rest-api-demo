protoc-generate:
	rm -rf pkg/api
	find api -iname "*.proto" | xargs -I@ echo --path=@ | xargs buf generate

statik-generate:
	statik -m -f -src third_party/openapi/

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		github.com/rakyll/statik \
		github.com/bufbuild/buf/cmd/buf


generate: protoc-generate statik-generate

