package main

import (
	"dlfile/dlstruct"
	"dlfile/domain/service"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", Hello)

	http.HandleFunc("/dlzip", DownZip)

	err := http.ListenAndServeTLS(":8080", "fugui_cert.pem", "fugui_key.pem", nil)
	if err != nil {
		// 发生异常打印日志
		fmt.Println(err, "监听端口错误")
	}

}

func Hello(w http.ResponseWriter, r *http.Request) {

	fmt.Println("23453534534")
	w.Write([]byte("hello"))
}

func DownZip(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/zip")

	// 创建数据文件夹
	dlstruct.FileN.SetDateFloderName()

	CreateFileDir(dlstruct.FileN.GetDateFloderName())

	// 创建zip文件夹
	dlstruct.FileN.SetZipDateFloderName()
	CreateFileDir(dlstruct.FileN.GetZipDateFloderName())

	fmt.Println("data文件夹: ", dlstruct.FileN.GetDateFloderName())
	fmt.Println("zip文件夹: ", dlstruct.FileN.GetZipDateFloderName())

	defer RemoveFloder()

	service.DownLoadZipSrv(w, r)

}

/**************************************************************************/

// 创建文件夹
func CreateFileDir(fileP string) {
	if !PathExists(fileP) {
		err := os.Mkdir(fileP, os.ModePerm)
		if err != nil {
			log.Printf("mkdir failed![%v]\n", err)
		} else {
			log.Printf("mkdir success!\n")
		}
	}
}

func PathExists(path string) bool {

	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 删除文件夹
func RemoveFloder() {

	os.RemoveAll(dlstruct.FileN.GetDateFloderName())
	os.RemoveAll(dlstruct.FileN.GetZipDateFloderName())

	fmt.Println("remove success")

}
