package main

import (
	"encoding/base64"
	"io/ioutil"
	"os"
)

func main(){
	//读原图片
	//ff, _ := os.Open("b.png")
	//defer ff.Close()
	//sourcebuffer := make([]byte, 500000)
	//n, _ := ff.Read(sourcebuffer)
	////base64压缩
	//sourcestring := base64.URLEncoding.EncodeToString(sourcebuffer[:n])
	////写入临时文件
	//ioutil.WriteFile("a.png.txt", []byte(sourcestring), 0667)
	////读取临时文件
	cc, _ := ioutil.ReadFile("a.png.txt")

	//解压
	dist, _ := base64.URLEncoding.DecodeString(string(cc))
	//写入新文件
	f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.Write(dist)

}

