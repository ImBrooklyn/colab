#!/bin/zsh

opt="paths=source_relative"
flag="require_unimplemented_servers=false"

function gen_code() {
    protoc --go_out="$1" --go_opt=$opt --go-grpc_out=$flag:"$1" --go-grpc_opt=$opt "$2"
}

gen_code ../common ./common/common.proto

gen_code ../grpc ./snowflake/snowflake_srv.proto

gen_code ../grpc ./message/message_srv.proto

gen_code ../grpc ./user/user_model.proto
gen_code ../grpc ./user/user_sign_srv.proto
gen_code ../grpc ./user/user_srv.proto
gen_code ../grpc ./user/organization_srv.proto

gen_code ../grpc ./portal/portal_menu_srv.proto

