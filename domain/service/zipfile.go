package service

import (
	"archive/zip"
	"bytes"
	"dlfile/dl"
	"dlfile/dlstruct"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func DownLoadZipSrv(w http.ResponseWriter, r *http.Request) {

	// 获取下载链接列表
	dl_urls := dl.GetFileListURL()

	ch := make(chan int)

	// 并发下载文件
	for _, v := range dl_urls {
		go dl.DownloadFile(v, ch)
	}

	for i := 0; i < len(dl_urls); i++ {
		// fmt.Println(dl_urls[i].FileId, " 下载完成")
		<-ch
	}

	// 压缩文件
	ZipWriter()

	// 读取压缩文件
	zipbyte2, err := ioutil.ReadFile(dlstruct.FileN.GetZipDateFileName())
	if err != nil {
		fmt.Println(err)
	}

	

	io.Copy(w, bytes.NewReader(zipbyte2))

}

// 压缩文件
func ZipWriter() {
	// 设置zip文件存放位置
	dlstruct.FileN.SetZipDateFileName()

	// 设置要压缩的文件夹
	baseFolder := dlstruct.FileN.GetDateFloderName() + "/"

	// Get a Buffer to Write To
	outFile, err := os.Create(dlstruct.FileN.GetZipDateFileName())
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, baseFolder, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		// fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

// fileinfo, err := ioutil.ReadDir("data")
// if err != nil {
// 	fmt.Println(err)
// }

// var zipbyte2 []byte
// for _, v := range fileinfo {
// 	zipbyte, err := ioutil.ReadFile("data/" + v.Name())
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	zipbyte2 = zipbyte

// }
