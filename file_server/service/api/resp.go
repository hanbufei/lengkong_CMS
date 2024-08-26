package api

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func ReturnSucess(ctx context.Context, res interface{}) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    200,
		Message: "",
		Data:    res,
	})
	return
}

func ReturnError(ctx context.Context, err error) {
	r := g.RequestFromCtx(ctx)
	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    500,
		Message: err.Error(),
		Data:    "",
	})
	return
}
