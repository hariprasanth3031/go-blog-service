go-blog-service

To generate go code using proto file:
protoc \                         
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
protos/blogs.proto

Db Confguration: \
1) Install mysql in local \
2) Create a database "blogs_management" \
3) Copy the schema from schema folder and create the tables \
4) Create a config.env file and fill the values for the keys in config/env.go file \

How to  check the response: \
1) Clone this repo \
2) Open Postman and create a new gRPC request \
3) Under service definiton tab, import the blogs.proto file from the repo \
4) Make the url as localhost:8080 \
5) Under message tab, click example message for a sample input \
6 )Click Invoke to see the response \

