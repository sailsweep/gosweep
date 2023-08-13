package notice

import "entdemo/svc"

type Notice struct{}

func NewNotice(svcCtx *svc.ServiceContext) (*Notice, *svc.ServiceContext) {
	svcCtx.InitDb(svcCtx.Config)
	return &Notice{}, svcCtx
}
