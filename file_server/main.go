package main

import (
	"context"
	"file_server/config"
	"file_server/service/controller"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"net/http"
)

var (
	cmd = gcmd.Command{
		Name:  "file_server",
		Usage: "file_server",
		Brief: "start file_server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetPort(config.Port)
			if config.SelfFileServer {
				s.AddStaticPath("/file", config.UploadPath)
				s.SetIndexFolder(true)
			}
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(Middleware) //默认的错误处理
				group.Group("/article", func(group *ghttp.RouterGroup) {
					group.Bind(
						controller.New(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)

func main() {
	cmd.Run(gctx.New())
}

func Middleware(r *ghttp.Request) {
	//校验http请求头内的key字段是否和配置的key一样
	if r.Request.Header.Get("key") != config.UploadKey {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    400,
			Message: "未获得授权",
			Data:    "",
		})
		return
	}
	MiddlewareCORS(r)
}

func MiddlewareCORS(r *ghttp.Request) {
	corsOptions := r.Response.DefaultCORSOptions()
	corsOptions.AllowDomain = config.CORS_allow
	r.Response.CORS(corsOptions)
	r.Middleware.Next()

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
	}

	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    code.Code(),
		Message: msg,
		Data:    res,
	})
}
