package dl

import (
	"dlfile/dlstruct"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"net/http"
)

func DownloadFile(file dlstruct.File, ch chan int) {

	// fmt.Println("开始下载：", file.FileId)

	resp, err := http.Get(file.DownloadUrl)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(dlstruct.FileN.GetDateFloderName() + "/img" + strconv.FormatInt(time.Now().UnixNano(), 10) + ".jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = f.Close() }()

	io.Copy(f, resp.Body)

	ch <- 1

}
