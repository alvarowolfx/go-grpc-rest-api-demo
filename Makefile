protoc-generate:
	find api -iname "*.proto" | xargs -I@ echo --path=@ | xargs buf generate


regenerate: protoc-generate

