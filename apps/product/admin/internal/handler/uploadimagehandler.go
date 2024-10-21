package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"webshop/apps/product/admin/internal/logic"
	"webshop/apps/product/admin/internal/svc"
)

func UploadImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewUploadImageLogic(r.Context(), svcCtx)
		resp, err := l.UploadImage()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
