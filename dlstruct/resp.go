package dlstruct

type FileList struct {

	// Intersection
	FileList []File `json:"file_list,omitempty"` // 文件列表

	// req
	AccessToken string `json:"access_token,omitempty"` // 接口调用凭证
	ENV         string `json:"env,omitempty"`          // 云环境ID

	// resp
	ErrCode int    `json:"errcode,omitempty"` // 错误码
	ErrMsg  string `json:"errmsg,omitempty"`  // 错误信息

	// other
	File     *File  `json:"file,omitempty"`     // 文件
	FileType string `json:"filetype,omitempty"` // 文件类型
}

type File struct {

	// Intersection
	FileId   string `json:"fileid,omitempty"`   // 文件ID
	FileName string `json:"filename,omitempty"` // 文件名

	// req
	MaxAge int `json:"max_age,omitempty"` // 下载链接有效期

	// resp
	DownloadUrl string `json:"download_url,omitempty"` // 下载链接
	Status      int    `json:"status,omitempty"`       // 状态码
	ErrMsg      string `json:"errmsg,omitempty"`       // 该文件错误信息
}
