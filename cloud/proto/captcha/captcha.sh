
cd /home/www/golang/src/apollo/cloud/proto;
protoc  --go-grpc_out=/home/www/golang/src/apollo/cloud/service  --go_out=/home/www/golang/src/apollo/cloud/service  captcha/captcha.proto;
protoc-go-inject-tag -input=/home/www/golang/src/apollo/cloud/service/captcha/captcha.pb.go;