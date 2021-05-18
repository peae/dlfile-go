package dlstruct

type UploadFile struct {

	// resp
	ErrCode       int    `json:"errcode,omitempty"`       //
	ErrMsg        string `json:"errmsg,omitempty"`        //
	URL           string `json:"url,omitempty"`           //
	Token         string `json:"token,omitempty"`         //
	Authorization string `json:"authorization,omitempty"` //
	FileId        string `json:"file_id,omitempty"`       //
	CosFileId     string `json:"cos_file_id,omitempty"`   //
}



