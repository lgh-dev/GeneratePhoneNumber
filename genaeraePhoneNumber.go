package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const path = "phone.txt"

func main() {
	err := ioutil.WriteFile(path, []byte{}, 0644)
	if err != nil {
		fmt.Printf("生成错误", err)
	}
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	defer f.Close()
	startTime := time.Now().Unix()
	var c int32 = 20000000
	var msg string
	n := int64(0)
	var i int32
	randNumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)
	for i = 1; i <= c; i++ {
		msg += "185" + strconv.Itoa(int(randNumber+i)) + "\n"
		//if c<1000{
		//var n int64
		//	n=int64((i-1)*12)
		//	_, err = f.WriteAt([]byte(msg), n)
		//}
		if i%100 == 0 {
			// 查找文件末尾的偏移量
			//var n int64
			//n=int64((i-1)*12)
			//n, _ := f.Seek(0, 2)
			_, err = f.WriteAt([]byte(msg), n)
			n += 1200
			msg = ""
		}
	}
	endTime := time.Now().Unix()
	fmt.Printf("完成耗时:%d秒", endTime-startTime)
	//err := ioutil.WriteFile(path, message, 0644)
	//if err != nil {
	//	fmt.Printf("生成错误",err)
	//}
}
