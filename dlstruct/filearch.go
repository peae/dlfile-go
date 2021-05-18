package dlstruct

import (
	"strconv"
	"time"
)

var FileN FileArch

const (
	baseDataDirName = "data"    //
	baseZipDirName  = "zipdata" //
	baseImgName     = "img"     //
	baseCsvName     = "csv"     //
	baseFileSuffix  = ".zip"    //
	baseJpgSuffix   = ".jpg"    //
	baseCsvSuffix   = ".csv"    //
)

type FileArch struct {
	FileName          string //
	FilePath          string //
	DataFloderName    string //
	DataFloderPath    string //
	ZipDataFloderName string //
	ZipDataFloderPath string //
	ZipFileName       string
	ZipFilePath       string
	FileSuffix        string //
}

// 设置数据文件夹
func (fileArch *FileArch) SetDateFloderName() {
	fileArch.DataFloderName = baseDataDirName + GetNowUnixNano()
}

// 获取数据文件夹
func (fileArch *FileArch) GetDateFloderName() string {
	return fileArch.DataFloderName
}

// 设置zip文件夹
func (fileArch *FileArch) SetZipDateFloderName() {
	fileArch.ZipDataFloderName = baseZipDirName + GetNowUnixNano()
}

// 获取zip文件夹
func (fileArch *FileArch) GetZipDateFloderName() string {
	return fileArch.ZipDataFloderName
}

// 设置zip文件路径
func (fileArch *FileArch) SetZipDateFileName() {
	fileArch.ZipFileName = fileArch.GetZipDateFloderName() + `/data` + GetNowUnixNano() + `.zip`
}

// 获取zip文件路径
func (fileArch *FileArch) GetZipDateFileName() string {
	return fileArch.ZipFileName
}

func (fileArch *FileArch) SetFileName() {
	fileArch.FileName = baseDataDirName + "/" + baseImgName + GetNowUnixNano() + fileArch.FileSuffix
}

func (fileArch *FileArch) GetFileName() string {
	return fileArch.FileName
}

func GetNowUnixNano() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
