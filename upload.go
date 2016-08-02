package main
//
////多文件上传
////作者：LvanNeo
////邮箱：lvan_software@foxmail.com
////日期：2014-04-17
////Golang + HTML5 实现多文件上传
//
//import (
//	"fmt"
//	"io"
//	"log"
//	"net/http"
//
//	"image/jpeg"
//
//	"image"
//)
//
//var UPLOAD="upload/"
////跳转上传页面
////作者：LvanNeo
////邮箱：lvan_software@foxmail.com
////日期：2014-04-17
//func indexHandle(w http.ResponseWriter, r *http.Request) {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("获取页面失败")
//		}
//	}()
//
//	// 上传页面
//	w.Header().Add("Content-Type", "text/html;charset=utf8")
//	w.WriteHeader(200)
//	w.Header().Add("Accept-Charset","uft8")
//	html := `
//		<html>
//	    <head>
//	        <title>Golang Upload Files</title>
//	    </head>
//	    <body>
//	        <form id="uploadForm"  enctype="multipart/form-data" action="/upload" method="POST">
//	            <p>Golang Upload</p> <br/>
//	            <input type="file" id="file1" name="userfile" multiple />	<br/>
//	            <input type="submit" value="Upload">
//	        </form>
//	   	</body>
//		</html>`
//	io.WriteString(w, html)
//}
//
////处理文件上传的 Web服务方法
////作者：LvanNeo
////邮箱：lvan_software@foxmail.com
////日期：2014-04-17
//func UploadServer(w http.ResponseWriter, r *http.Request) {
//	defer func() {
//		if err := recover(); err != nil {
//			fmt.Println("文件上传异常")
//			panic(err)
//		}
//	}()
//
//	if "POST" == r.Method {
//
//		r.ParseMultipartForm(32 << 20)	//在使用r.MultipartForm前必须先调用ParseMultipartForm方法，参数为最大缓存
//
//
//
//		if r.MultipartForm != nil && r.MultipartForm.File != nil {
//			fhs := r.MultipartForm.File["userfile"]		//获取所有上传文件信息
//			num := len(fhs)
//			list:=make([]image.Image,0)
//			fmt.Println("总文件数：%d 个文件", num)
//
//			//region for
//			for _, fheader := range fhs {//循环对每个文件进行处理
//
//				//获取文件名
//				filename := fheader.Filename
//
//				//结束文件
//				file,err := fheader.Open()
//				if err != nil {
//					fmt.Println(err)
//				}
//				defer file.Close()
//				img,err:=jpeg.Decode(file)
//				if(err!=nil){
//					panic(err)
//				}
//				list=append(list,img)
//				fmt.Println(filename,img.Bounds())
//			}
//			//endregion
//
//
//		}
//
//		return
//	} else {
//		indexHandle(w,r)
//	}
//
//}
//
//func main() {
//	fmt.Println("Listening Port: 8086")
//	http.HandleFunc("/", indexHandle)
//	http.HandleFunc("/upload", UploadServer)
//	err := http.ListenAndServe(":8086", nil)
//	if err != nil {
//		log.Fatal("ListenAndServe: ", err)
//	}
}