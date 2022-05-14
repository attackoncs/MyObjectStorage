使用go开发分布式对象存储系统
# 命令
## 运行程序
go run server.go
## 创建objects下的test1文件
mkdir ~/objects

vim test1
## 使用curl模拟get和put请求
curl 127.0.0.1:1280/objects/test1

curl -v 127.0.0.1:1280/objects/test1 -XPUT -d "put test"
