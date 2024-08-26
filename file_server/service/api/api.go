package api

import "github.com/gogf/gf/v2/frame/g"

type DefaultRes struct{}

// 返回随机目录
type GetPathReq struct {
	g.Meta `path:"/getPath" method:"get" summary:"获取随机目录"`
}

// 无特定返回
type FileUploadReq struct {
	g.Meta `path:"/upload" method:"post" summary:"文件上传"`
	File   string `json:"file" type:"file" dc:"选择上传文件"`
	Path   string `p:"path" v:"required#路径不能为空"`
}

// 返回生成的图片路径
type ImgUploadReq struct {
	g.Meta `path:"/uploadImg" method:"post" summary:"图片上传"`
	Img    string `json:"img" type:"file" dc:"选择上传图片"`
	Path   string `p:"path" v:"required#路径不能为空"`
}

// 无特定返回
type DeleteReq struct {
	g.Meta `path:"/delete" method:"get" summary:"删除文件"`
	Path   string `p:"path" v:"required#路径不能为空"`
}

//
//type FileUploadRes struct {
//	Result int `json:"result"`
//}
//
//type ImgUploadReq struct {
//	g.Meta   `path:"/uploadImg" method:"post" summary:"图片上传"`
//	Img string `json:"img" type:"file" dc:"选择上传图片"`
//	DirName string `p:"dirName" v:"required#Key不能为空"`
//	Key string `p:"key" in:"header" v:"required#Key不能为空"`
//}
//
//type ImgUploadRes struct {
//	ImgName string `json:"imgName"`
//}
//
//type ClearHistoryReq struct {
//	g.Meta   `path:"/clearHistory" method:"get" summary:"清除历史不用文章"`
//	Key string `p:"key"  v:"required#Key不能为空"`
//}
//
//type ClearHistoryRes struct {
//	Result int `json:"result"`
//}
//
//type DeleteFileReq struct {
//	g.Meta   `path:"/delete" method:"get" summary:"删除文件"`
//	Path string `p:"path"  v:"required#Path不能为空"`
//	Key string `p:"key"  v:"required#Key不能为空"`
//}
//
//type DeleteFileRes struct {
//	Result int `json:"result"`
//}
