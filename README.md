# Prerequisites

* Install protobuf 3
* Install glide 
```
mkdir bin
curl https://glide.sh/get | sh
```

* Make sure have include/google folder copied to /usr/local/include
* Make sure to install all dependencies, this will add submodule folders under src/.. you can delete them

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
```

# Generating api proto

protoc -I /usr/local/include \
 -I $GOPATH/resources \
 -I $GOPATH/src/api/vendor/google.golang.org \
 -I $GOPATH/src/api/vendor/github.com/golang/protobuf \
 -I $GOPATH/src/api/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 $GOPATH/resources/proto/api.proto \
 --go_out=plugins=grpc:$GOPATH/src/api/api_proto

# Generating gateway reverse proxy

protoc -I /usr/local/include \
 -I $GOPATH/resources \
 -I $GOPATH/src/api/vendor/google.golang.org \
 -I $GOPATH/src/api/vendor/github.com/golang/protobuf \
 -I $GOPATH/src/api/vendor/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:$GOPATH/src/api/api_proto \
 $GOPATH/resources/proto/api.proto 


