* b站教学视频 https://www.bilibili.com/video/BV1Ev411w7yu
* https://www.bilibili.com/video/BV1D34y1Q7aH

* (安装protoc) brew install protobuf
* go install goo两个install(官网)
* (在项目中)执行(可以发现go.mod文件的变化，添加依赖) go get google.golang.org/grpc 
* 进入pb(proto文件所在)文件夹执行(只需要把最后的文件名修改即可) protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  hello_grpc.proto
* 代码就在pb文件夹中生成
