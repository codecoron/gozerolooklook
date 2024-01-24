package upload

import (
	uuid "github.com/satori/go.uuid"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io/ioutil"
	"looklook/app/upload/cmd/api/internal/logic/upload"
	"looklook/app/upload/cmd/api/internal/svc"
	"looklook/app/upload/cmd/api/internal/types"
	"looklook/common/result"
	"net/http"
	"path/filepath"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUploadReq
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
			return
		}
		file, header, err := r.FormFile("file")
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
			return
		}

		//获取文件类型
		ext := filepath.Ext(header.Filename)
		//获取文件内容
		all, err := ioutil.ReadAll(file)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
			return
		}
		//填充参数
		//req.FileName = header.Filename
		//使用uuid生成文件名
		req.FileName = uuid.NewV4().String()
		req.Size = header.Size
		req.Ext = ext
		req.FileData = all

		l := upload.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(&req)
		if err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			result.ParamErrorResult(r, w, err)
		} else {
			//httpx.OkJsonCtx(r.Context(), w, resp)
			result.HttpResult(r, w, resp, err)
		}
	}
}
