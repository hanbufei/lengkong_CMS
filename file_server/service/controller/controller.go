package controller

import (
	"context"
	"file_server/config"
	"file_server/service/api"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
	"io"
	"strconv"
	"strings"
)

var Article = Controller{}

type Controller struct {
}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) GetPath(ctx context.Context, req *api.GetPathReq) (res *api.DefaultRes, err error) {
	data := gtime.Date()
	dirName := data + "_" + strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(6))
	api.ReturnSucess(ctx, dirName)
	return
}

func (c *Controller) Upload(ctx context.Context, req *api.FileUploadReq) (res *api.DefaultRes, err error) {
	if req.File == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择文件")
	}
	if strings.Contains(req.Path, "/") || strings.Contains(req.Path, "\\") || strings.Contains(req.Path, "..") {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "路径含不被允许字符"))
		return
	}
	fileName := req.Path + gfile.Separator + "content.json"
	err = gfile.PutContents(config.UploadPath+gfile.Separator+fileName, req.File)
	if err != nil {
		api.ReturnError(ctx, err)
		return
	}
	api.ReturnSucess(ctx, "上传成功")
	return
}

func (c *Controller) ImgUpload(ctx context.Context, req *api.ImgUploadReq) (res *api.DefaultRes, err error) {
	if req.Img == "" {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请选择文件")
	}
	if strings.Contains(req.Path, "/") || strings.Contains(req.Path, "\\") || strings.Contains(req.Path, "..") {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "路径含不被允许字符"))
		return
	}
	name := strings.ToLower(strconv.FormatInt(gtime.TimestampNano(), 36)+grand.S(2)) + ".png"
	f := g.RequestFromCtx(ctx).GetUploadFile("img")
	file, err := f.Open()
	if err != nil {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "图片解析失败"))
		return
	}
	defer file.Close()
	imgPath := gfile.Join(config.UploadPath, req.Path, name)
	newImg, err := gfile.Create(imgPath)
	if err != nil {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "图片解析失败"))
		return
	}
	defer newImg.Close()
	if _, err = io.Copy(newImg, file); err != nil {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "图片解析失败"))
		return
	}
	api.ReturnSucess(ctx, gfile.Join(req.Path, name))
	return
}

func (c *Controller) Delete(ctx context.Context, req *api.DeleteReq) (res *api.DefaultRes, err error) {
	if strings.Contains(req.Path, "/") || strings.Contains(req.Path, "\\") || strings.Contains(req.Path, "..") {
		api.ReturnError(ctx, gerror.NewCode(gcode.CodeInternalError, "路径含不被允许字符"))
		return
	}
	err = gfile.Remove(config.UploadPath + gfile.Separator + req.Path)
	if err != nil {
		api.ReturnError(ctx, err)
		return
	}
	api.ReturnSucess(ctx, "删除成功")
	return
}
