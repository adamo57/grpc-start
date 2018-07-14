 # gRPC Code Generation Commands
 
 ### Generate New Proto File
 ``` 
 protoc -I user/ user/userService.proto --go_out=plugins=grpc:user
 ```

 ### Generate gRPC Stub
 ```
 protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --go_out=plugins=grpc:. \
  path/to/your_service.proto
  ```

 ### Generate Reverse Proxy:  

``` 
protoc -I /usr/local/include -I. \
 -I $HOME/go/src \
 -I $HOME/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
 --grpc-gateway_out=logtostderr=true:. \
 user/userService.proto