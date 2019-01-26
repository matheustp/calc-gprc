 protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --plugin=protoc-gen-swagger=$GOPATH/bin/protoc-gen-swagger \
  --swagger_out=logtostderr=true:. \
  pb/*.proto