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
		fmt.Printf("生成错误",err)
	}
	f, err := os.OpenFile(path, os.O_WRONLY, 0644)
	defer f.Close()
	startTime:= time.Now().Unix();
	for i:=0;i<2000000 ;i++  {
		randNumber:= rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000)
		msg:="185"+strconv.Itoa(int(randNumber))+"\n";

		// 查找文件末尾的偏移量
			n, _ := f.Seek(0, 2)
			_, err = f.WriteAt([]byte(msg), n)
	}
	endTime:=time.Now().Unix();
	fmt.Sprintf("完成耗时:{}毫秒",endTime-startTime)
	//err := ioutil.WriteFile(path, message, 0644)
	//if err != nil {
	//	fmt.Printf("生成错误",err)
	//}
}
