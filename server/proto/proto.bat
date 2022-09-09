set PBGO_OUT_DIR=../gen/go
protoc --go_out %PBGO_OUT_DIR% --go_opt paths=source_relative --go-grpc_out %PBGO_OUT_DIR% --go-grpc_opt paths=source_relative trip/trip.proto
protoc --grpc-gateway_out %PBGO_OUT_DIR% --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative trip/trip.proto

set PBTS_BIN_DIR=..\..\wx\miniprogram\node_modules\.bin
set PBTS_OUT_DIR=..\..\wx\miniprogram\gen\ts
%PBTS_BIN_DIR%\pbjs -t static -w es6 .\trip\trip.proto --no-create --no-encode --no-verify --no-delimited -o %PBTS_OUT_DIR%\trip_pb.js
%PBTS_BIN_DIR%\pbts -o %PBTS_OUT_DIR%\trip_pb.d.ts %PBTS_OUT_DIR%\trip_pb.js