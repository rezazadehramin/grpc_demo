go-proto:
	protoc --go-grpc_out=./protos \
       --go_out=./protos \
       ./protos/everphone.proto

php-proto:
	protoc --php_out=./protos/php \
	  	--grpc_out=./protos/php \
	  	--proto_path=./protos \
       	--plugin=protoc-gen-grpc=/grpc/cmake/build/grpc_php_plugin \
       ./protos/everphone.proto
