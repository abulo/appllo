
cd /home/www/golang/src/apollo/cloud/proto/system;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service/system  --go_out=/home/www/golang/src/apollo/cloud/service/system  tenant/system_tenant.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/system/tenant/system_tenant.pb.go;