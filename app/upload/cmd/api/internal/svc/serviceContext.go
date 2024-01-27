package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"looklook/app/upload/cmd/api/internal/config"
	"looklook/app/upload/cmd/rpc/upload"
)

type ServiceContext struct {
	Config        config.Config
	FileUploadRpc upload.Upload
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		FileUploadRpc: upload.NewUpload(zrpc.MustNewClient(c.FileUploadRpcConf)),
	}
}
