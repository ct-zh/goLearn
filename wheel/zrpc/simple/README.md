# 使用zrpc
1. 创建proto文件：`hello.proto`;
2. 生成pb文件：`protoc --go_out=. --go--grpc_out=. hello.proto`; 
3. 编写server端，注册etcd;
4. 编写client端，