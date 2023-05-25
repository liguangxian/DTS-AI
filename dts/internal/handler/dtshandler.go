package handler

import (
	"net/http"

	"DTS-AI/dts/internal/logic"
	"DTS-AI/dts/internal/svc"
	"DTS-AI/dts/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DtsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDtsLogic(r.Context(), svcCtx)
		resp, err := l.Dts(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
