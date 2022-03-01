#!/bin/sh

SRC_DIR=.
DST_DIR=../../../../../../..

GOGOPROTO_PATH=$GOPATH/src/github.com/gogo/protobuf/protobuf
MTPROTO_PATH=$GOPATH/src/github.com/teamgram/proto/mtproto

protoc -I=$SRC_DIR:$MTPROTO_PATH --proto_path=$GOPATH/src:$GOGOPROTO_PATH:./ \
    --gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,:$DST_DIR \
    $SRC_DIR/*.proto
