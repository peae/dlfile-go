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

// {
// 	"errcode": 0,
// 	"errmsg": "ok",
// 	"url": "https://cos.ap-shanghai.myqcloud.com/636c-cloud1-1ger9dnb820d9272-1305935445/dlzip.zip",
// 	"token": "VSeo4eBJkkCnEvBtnQCJwlj1RdduPANa628e50959a91462604acbcff69dbeb92PTRDvluRds74sax4RAWjj4VEFj-HzXgtbSmVeKFdT-MutQolVpyAY-1Jm89MfiRVmwO6WdvTyGP2cvfwxscC4N5-WoLufGKHALZdc3estXuNHUYK-RGA5hRWaIE_Ig4JCZlXwoCPPYl-m5-y_Jq3QyttkoUmXcnZ6__72NHuPPkHN9ETqu8Id9SeRc-0aProAa6d8M-nSWzEQQVUrX_s_ApWxrwLsxr_XHvRa1HdCjDOJYJ8_9N1eHLCS1Kr1CqLYwRzUUrXeBdXhVo6SelH-emMDlnOMMxS--mqOnyFyRZNjKqBL5RTjMuRjCn2BLbDhqzTrLjy5Gf7BCfvm53FoV8EWSUEa4ntX6rBnc0XZQ0",
// 	"authorization": "q-sign-algorithm=sha1&q-ak=AKIDGKwPgdL59dMdPGagP7oE9xrgJYf-3wtPnEpAPO3nnNWc-lgT1izzXdxj1W7DQTA9&q-sign-time=1621333784;1621334684&q-key-time=1621333784;1621334684&q-header-list=&q-url-param-list=&q-signature=28514209c941e09d94824e3b460f93cff987ef2d",
// 	"file_id": "cloud://cloud1-1ger9dnb820d9272.636c-cloud1-1ger9dnb820d9272-1305935445/dlzip.zip",
// 	"cos_file_id": "HKknVK2y5ZbocHzuNnn+IPuqac68wfh0DU8FlEU4Uho44wpsi9807b69KdvGbKwzZtdv/JeUGd14Z4i374ljIX0VWUlP47QnCZMRi9LakU+ZU46TsXeifRbJfzl484aCyZH95FTYBoBY08XxJ0Wt4W5DWbj0hHY8dRx1PEGzWZLKqJ8="
// }
