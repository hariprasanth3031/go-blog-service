go-blog-service

To generate go code using proto file:
protoc \                         
--go_out=. \
--go_opt=paths=source_relative \
--go-grpc_out=. \
--go-grpc_opt=paths=source_relative \
protos/blogs.proto

Db Confguration:
Install mysql in local
Create a database "blogs_management"
Copy the schema from schema folder and create the tables
Create a config.env file and fill the values for the keys in config/env.go file

How to  check the response:
Clone this repo
Open Postman and create a new gRPC request
Under service definiton tab, import the blogs.proto file from the repo
Make the url as localhost:8080 
Under message tab, click example message for a sample input
Click Invoke to see the response

