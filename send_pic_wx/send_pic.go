package main;

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/",welcome)
	http.HandleFunc("/upload", uploadHandle) // 上传
	http.HandleFunc("/uploaded/", showPicHandle)  //显示图片
	err := http.ListenAndServe(":8011", nil)
	fmt.Println(err)
}
func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("1.html")
	t.Execute(w, nil)
}
// 上传图像接口
func uploadHandle (w http.ResponseWriter, req *http.Request) {
	//fmt.Println("gggg")

	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
		fmt.Println("get")

	} else {
		// 接收图片
		fmt.Println("gggg")

		s := req.FormValue("data")
		//errorHandle(err, w)
		//fmt.Println(s)
		//fmt.Println(handle)
		//cc,_:=base64.StdEncoding.DecodeString(s)
		//fmt.Println(sd)
		dist, _ := base64.URLEncoding.DecodeString(string(s))
		f, _ := os.OpenFile("xx.png", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(dist)
		//// 检查图片后缀
		//ext := strings.ToLower(path.Ext(handle.Filename))
		//if ext != ".jpg" && ext != ".png" {
		//	errorHandle(errors.New("只支持jpg/png图片上传"), w);
		//	return
		//	//defer os.Exit(2)
		//}
		//
		//// 保存图片
		//os.Mkdir("./uploaded/", 0777)
		//saveFile, err := os.OpenFile("./uploaded/" + handle.Filename, os.O_WRONLY|os.O_CREATE, 0666);
		//errorHandle(err, w)
		//io.Copy(saveFile, uploadFile);
		//
		//defer uploadFile.Close()
		//defer saveFile.Close()
		//// 上传图片成功
		//w.Write([]byte("查看上传图片: <a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a>"));
	}
}

// 显示图片接口
func showPicHandle( w http.ResponseWriter, req *http.Request ) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w);

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w);
	w.Write(buff)
}

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if  err != nil {
		w.Write([]byte(err.Error()))
	}
}

const html = `<html>
    <head></head>
    <body>
        <form method="post" enctype="multipart/form-data">
            <input type="file" name="image" />
            <input type="submit" />
        </form>
    </body>
</html>`