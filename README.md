go-blog-service

To generate go code using proto file:
protoc \                         
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
protos/blogs.proto

How to Run:

Clone this repo
Open Postman and create a new gRPC request
Under service definiton tab, import the blogs.proto file from the repo
Make the url as localhost:8080 
Under message tab, click example message for a sample input
Click Invoke to see the response
