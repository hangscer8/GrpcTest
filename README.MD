* b站教学视频 https://www.bilibili.com/video/BV1Ev411w7yu
* https://www.bilibili.com/video/BV1D34y1Q7aH
* 官网文档

* (安装protoc) brew install protobuf
* go install goo两个install(官网)
* (在项目中)执行(可以发现go.mod文件的变化，添加依赖) go get google.golang.org/grpc
* 进入pb(proto文件所在)文件夹执行(只需要把最后的文件名修改即可)
  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative hello_grpc.proto
* 代码就在pb文件夹中生成

代码逻辑

* 服务端
    * 取出server(继承?UnimplementedHelloGRPCServer，为其实现?方法)
    * 挂载方法
    * 注册服务
    * 创建监听

* 客户端
    * 创建连接
    * new client
    * 调用client方法
    * 返回值

