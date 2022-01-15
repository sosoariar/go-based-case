#基础数据类型
同一个作用域中不允许同名变量

Terminal 运行方式
```shell
# 源代码文件目录下
go run main.go

# 编译一个可执行文件到目录下
go build main.go
# 直接执行exe文件
./main.exe
```

自带的格式化代码指令
```shell
go fmt main.go
```
常量

iota 从赋值开始,遇到常量就计数,遇到 const 关键字就重置
iota 遇到不需要计数的,使用_替换