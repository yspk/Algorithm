###答案演示
####环境要求
+ macOS
+ golang
+ clang
+ shell
####演示步骤
1. 安装pcre2
```shell
brew install pcre2
```
2. 启动udp server,监听执行结果
```shell
chmod +x shellserver.sh
./shellserver.sh
```
3. 执行主程序
```shell
go run main.go
```
