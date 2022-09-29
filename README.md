# gRPC Unary 

```
protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
book/book.proto
```

## Penjelasan

Unary gRPC adalah komunikasi client dan server dimana mengirim satu request dan menerima satu response.

![image](https://storage.googleapis.com/kotakode-prod-public/images/55e9fda8-8692-4cf3-903d-0b92a0edc9a7-unary_grpc.png)


