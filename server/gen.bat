set PBGO_OUT_DIR=.
@REM paths=source_relative 这个设置了的话，生成的文件和 .proto 文件会放在同一目录下
protoc --go_out %PBGO_OUT_DIR% --go_opt paths=source_relative --go-grpc_out %PBGO_OUT_DIR% --go-grpc_opt paths=source_relative auth/api/gen/v1/auth.proto
protoc --grpc-gateway_out %PBGO_OUT_DIR% --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative auth/api/gen/v1/auth.proto

protoc --go_out %PBGO_OUT_DIR% --go_opt paths=source_relative --go-grpc_out %PBGO_OUT_DIR% --go-grpc_opt paths=source_relative rental/api/gen/v1/rental.proto
protoc --grpc-gateway_out %PBGO_OUT_DIR% --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative rental/api/gen/v1/rental.proto

protoc --go_out %PBGO_OUT_DIR% --go_opt paths=source_relative --go-grpc_out %PBGO_OUT_DIR% --go-grpc_opt paths=source_relative blob/api/gen/v1/blob.proto

protoc --go_out %PBGO_OUT_DIR% --go_opt paths=source_relative --go-grpc_out %PBGO_OUT_DIR% --go-grpc_opt paths=source_relative car/api/gen/v1/car.proto
protoc --grpc-gateway_out %PBGO_OUT_DIR% --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative car/api/gen/v1/car.proto

set PBTS_BIN_DIR=..\wx\miniprogram\node_modules\.bin
set PBTS_OUT_DIR=..\wx\miniprogram\gen\ts\auth
..\wx\miniprogram\node_modules\.bin\pbjs -t static -w es6 car\api\gen\v1\car.proto --no-create --no-encode --no-verify --no-delimited -o %PBTS_OUT_DIR%\car_pb.js
..\wx\miniprogram\node_modules\.bin\pbts -o %PBTS_OUT_DIR%\car_pb.d.ts %PBTS_OUT_DIR%\car_pb.js

..\wx\miniprogram\node_modules\.bin\pbjs -t static -w es6 rental\api\gen\v1\rental.proto --no-create --no-encode --no-verify --no-delimited -o %PBTS_OUT_DIR%\rental_pb.js
..\wx\miniprogram\node_modules\.bin\pbts -o %PBTS_OUT_DIR%\rental_pb.d.ts %PBTS_OUT_DIR%\rental_pb.js

..\wx\miniprogram\node_modules\.bin\pbjs -t static -w es6 auth\api\gen\v1\auth.proto --no-create --no-encode --no-verify --no-delimited -o %PBTS_OUT_DIR%\auth_pb.js
..\wx\miniprogram\node_modules\.bin\pbts -o %PBTS_OUT_DIR%\auth_pb.d.ts %PBTS_OUT_DIR%\auth_pb.js