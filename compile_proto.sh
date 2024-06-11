# https://protobuf.dev/getting-started/gotutorial/#compiling-protocol-buffers
protoc ./rin.proto --go_out=./pkg/ --go-grpc_out=./pkg/
