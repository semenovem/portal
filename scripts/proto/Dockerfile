FROM curlimages/curl:7.84.0 as builder

WORKDIR /tmp
ARG url_protoc="https://github.com/protocolbuffers/protobuf/releases/download/v21.4/protoc-21.4-linux-x86_64.zip"
ARG url_buf="https://github.com/bufbuild/buf/releases/download/v1.9.0/buf-Linux-x86_64.tar.gz"

RUN curl -o "protoc.zip" -L "${url_protoc}"
RUN unzip "protoc.zip"

RUN curl -o "buf.tar.gz" -L "${url_buf}"
RUN tar -zxf buf.tar.gz


FROM golang:1.18.5
WORKDIR /app
COPY --from=0 /tmp/bin/protoc /go/bin
COPY --from=0 /tmp/include /go/bin
COPY --from=0 /tmp/buf /opt/buf

ENV BUF_CACHE_DIR=/tmp/caches
ENV PATH="${PATH}:/opt/buf/bin"

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0 && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.14.0
