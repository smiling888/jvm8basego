package main

import "flag"
import "fmt"
import "os"

//在Java语言中，API一般以类库的形式提供。在Go语言中，API 则是以包(package)的形式提供。包可以向用户提供常量、变量、结 构体以及函数等
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string //这里指定jre目录的位置
	class       string
	args        []string
}

//1 os包定义了一个Args变量，其中存放传递给命令行的全部参数。如果直接处理os.Args变量，需要写很多代码。
//2 Go语言内置了flag包，这个包可以帮助我们处理命令行选项。
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
