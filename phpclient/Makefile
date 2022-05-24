proto:
	protoc --php_out=./protos \
	  	--grpc_out=./protos \
	  	--proto_path=./protos \
       	--plugin=protoc-gen-grpc=/grpc/cmake/build/grpc_php_plugin \
       ./protos/everphone.proto