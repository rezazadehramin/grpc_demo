FROM php:7.4-cli

# Environment variables
ENV DEBIAN_FRONTEND=noninteractive

# Utils
RUN apt-get update -yqq && \
    apt-get upgrade -yqq && \
    apt-get install -y unzip build-essential git software-properties-common curl pkg-config zip zlib1g-dev

# Composer installation
COPY --from=composer:latest /usr/bin/composer /usr/local/bin/composer

# Install grpc and probuf with pecl
RUN pecl install grpc && pecl install protobuf

# Enable grpc and protobuf extensions in php.ini file
RUN echo starting && \
    docker-php-ext-enable grpc && \
    docker-php-ext-enable protobuf

# Install cmake
RUN apt-get update -yqq && apt-get -y install cmake

# Install grpc_php_plugin and protoc
RUN git clone -b v1.36.2 https://github.com/grpc/grpc && \
    cd grpc && git submodule update --init && \
    mkdir cmake/build && cd cmake/build && \
    cmake ../.. && make protoc grpc_php_plugin

# Setting node, protoc and grpc_php_plugin paths
ENV PATH "/grpc/cmake/build:${PATH}"
ENV PATH "/grpc/cmake/build/third_party/protobuf:${PATH}"


# Moving client folder to vm
WORKDIR root/app/
COPY ./protos ./protos
COPY Makefile .

# Packages
#RUN composer install

# Start the client app
#ENTRYPOINT  [ "make", "proto" ]