package config

var (
	Port       = 9999                               //服务运行端口
	UploadPath = "/Users/hanfei/Desktop/upload"     //本地文件存储位置
	CORS_allow = []string{"localhost", "127.0.0.1"} //允许cors的白名单

	UploadKey      = "d282a7vrhcq0kvzp39" //请求头key字段，用于API服务的访问控制
	SelfFileServer = true                 //提供file路径用于加载文件读取服务。建议采用nginx并改为false
)
