protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --plugin=protoc-gen-grpc-gateway=$GOPATH/bin/protoc-gen-grpc-gateway \
  --grpc-gateway_out=logtostderr=true:. \
  pb/*.proto