package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

// 系统访问日志
type AccessLoggerFormat struct {
	IP        string `json:"IP"`
	Header    string `json:"Header"`
	UserAgent string `json:"UserAgent"`
	Extend    string `json:"Extend"`
}

// 系统正常打印日志
type InfoLoggerFormat struct {
	Package string `json:"Package"`
	Method  string `json:"Method"`
	Message string `json:"Message"`
}

// 系统异常日志
type ErrorLoggerFormat struct {
	Package      string `json:"Package"`
	Method       string `json:"Method"`
	ErrorMessage string `json:"ErrorMessage"`
}

var AccessPath = "/Users/artorias/Desktop/logs/access"
var InfoPath = "/Users/artorias/Desktop/logs/info"
var ErrorPath = "/Users/artorias/Desktop/logs/error"

func GetLogger(logPath, logType string) *log.Logger{
	// 打开日志文件
	// 第二个参数为打开文件的模式，可选如下：
	/*
		O_RDONLY // 只读模式打开文件
		O_WRONLY // 只写模式打开文件
		O_RDWR   // 读写模式打开文件
		O_APPEND // 写操作时将数据附加到文件尾部
		O_CREATE // 如果不存在将创建一个新文件
		O_EXCL   // 和O_CREATE配合使用，文件必须不存在
		O_SYNC   // 打开文件用于同步I/O
		O_TRUNC  // 如果可能，打开时清空文件
	 */

	// 第三个参数为文件权限，请参考linux文件权限，664在这里为八进制，代表：rw-rw-r--
	logFileName := "/log_" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(logPath+logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// 第一个参数为输出io，可以是文件也可以是实现了该接口的对象，此处为日志文件；第二个参数为自定义前缀；第三个参数为输出日志的格式选项，可多选组合
	// 第三个参数可选如下：
	/*
		Ldate         = 1             // 日期：2009/01/23
		Ltime         = 2             // 时间：01:23:23
		Lmicroseconds = 4             // 微秒分辨率：01:23:23.123123（用于增强Ltime位）
		Llongfile     = 8             // 文件全路径名+行号： /a/b/c/d.go:23
		Lshortfile    = 16            // 文件无路径名+行号：d.go:23（会覆盖掉Llongfile）
		LstdFlags     = Ldate | Ltime // 标准logger的初始值
	 */
	logger := log.New(logFile, "["+logType+"]", log.Ldate|log.Ltime|log.Llongfile)

	// 日志输出
	//logger.Print("日志测试Print输出，处理同fmt.Print")
	//logger.Println("日志测试Println输出，处理同fmt.Println")
	//logger.Printf("日志测试%s输出，处理同fmt.Printf", "Printf")
	//
	//// 日志输出，同时直接终止程序，后续的操作都不会执行
	//logger.Fatal("日志测试Fatal输出，处理等价于：debugLog.Print()后，再执行os.Exit(1)")
	//logger.Fatalln("日志测试Fatalln输出，处理等价于：debugLog.Println()后，再执行os.Exit(1)")
	//logger.Fatalf("日志测试%s输出，处理等价于：debugLog.Print()后，再执行os.Exit(1)", "Fatalf")
	//
	// 日志输出，同时抛出异常，可用recover捕捉
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("===========", r)
		}
	}()
	//logger.Panic("日志测试Panic输出，处理等价于：debugLog.Print()后，再执行Panic()")
	//logger.Panicln("日志测试Panicln输出，处理等价于：debugLog.Println()后，再执行Panic()")
	//logger.Panicf("日志测试%s输出，处理等价于：debugLog.Printf()后，再执行Panic()", "Panicf")
	//
	//fmt.Println("前缀为：", debugLog.Prefix())    // 前缀为： [debug]
	//fmt.Println("输出选项为：", debugLog.Flags()) // 输出选项为： 11
	//// 设置前缀
	//logger.SetPrefix("[info]")
	//// 设置输出选项
	//logger.SetFlags(log.LstdFlags)
	//fmt.Println("前缀为：", debugLog.Prefix())    // 前缀为： [info]
	//fmt.Println("输出选项为：", debugLog.Flags()) // 输出选项为： 3
	return logger
}

func GetConsoleLogger() *log.Logger{
	logger := log.New(os.Stdout, "[Debug]", log.Ldate|log.Ltime|log.Llongfile)
	return logger
}