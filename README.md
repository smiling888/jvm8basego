# 列表

# 常用命令
```
 go mod init src/demo
 
 go build -o osArgsDemo  不用带后缀
 
 go run *.go -cp filelzh/zzz  MyApp  123 144
 
 报错：package Projectname is not in GOROOT (C:\Program Files\Go\src\Projectname)
 go env -w GO111MODULE=off


```

- go build：用于测试编译包，在项目目录下生成可执行文件（有main包）。

- go install：主要用来生成库和工具。一是编译包文件（无main包），将编译后的包文件放到 pkg 目录下（$GOPATH/pkg）。二是编译生成可执行文件（有main包），将可执行文件放到 bin 目录（$GOPATH/bin）。

- 运行ch01下对main时, 可以同时选额多个文件运行 
# 日志
20210317 读取class文件
./ch02/main -Xjre "/Library/Java/JavaVirtualMachines/jdk1.8.0_71.jdk/Contents/Home/jre" java.lang.Object

20210307 创建项目
