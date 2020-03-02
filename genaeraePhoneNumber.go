package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const path = "phone.txt"

var phoneNumber = 20000000 //号码数
var goNumber = 100         //协程数

func main() {
	startTime := time.Now().UnixNano()
	f := getFile()  //获取文件连接
	defer f.Close() //关闭连接
	//write1(f) //单线程一个个写
	//write2(f) //多协程写入
	write3(f) //多协程批量写
	endTime := time.Now().UnixNano()
	fmt.Printf("完成耗时:%d毫秒", (endTime-startTime)/1000000)
}

//获取文件连接
func getFile() *os.File {
	err := ioutil.WriteFile(path, []byte{}, 0644)
	if err != nil {
		fmt.Printf("创建文件异常", err)
	}
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("打开文件连接异常", err)
	}
	return f
}

//单线程一个个写
func write1(f *os.File) {
	var msg string
	n := 0 //文件末尾的偏移量
	for i := 1; i <= phoneNumber; i++ {
		msg = "185" + strconv.Itoa(100000000+i) + "\n"
		n = (i - 1) * 12 //计算偏移量
		f.WriteAt([]byte(msg), int64(n))
		msg = ""
	}
}

//单线程批量写
func write2(f *os.File) {
	var msg string
	n := 0 //文件末尾的偏移量
	for i := 1; i <= phoneNumber; i++ {
		msg = "185" + strconv.Itoa(100000000+i) + "\n"
		if i%20 == 0 { // 批量写入
			n = (i - 1) * 12 //计算偏移量
			f.WriteAt([]byte(msg), int64(n))
			msg = ""
		}
	}
}

//多协程批量写入。
func write3(f *os.File) {
	ch := make(chan int, goNumber) // 通道
	dur := phoneNumber / goNumber  // 任务间隔数
	for startGo := 0; startGo < goNumber; startGo++ {
		begin := startGo * dur     //任务分片起点
		end := (startGo + 1) * dur //任务分片终点
		// 写入方法
		writeFunc := func(begin, end int) {
			var msg string
			n := 0 //文件末尾的偏移量
			for i := begin; i <= end; i++ {
				msg += "185" + strconv.Itoa(100000000+i) + "\n"
				if i%20 == 0 {
					n = (i - 1) * 12 //计算偏移量
					f.WriteAt([]byte(msg), int64(n))
					msg = ""
				}
			}
			ch <- 1
		}
		go writeFunc(begin, end)
	}
	//等待协程执行完毕。
	for {
		if goNumber == 0 {
			break
		}
		goNumber -= <-ch
		fmt.Printf("go number:%d\n", goNumber)
	}
}
