package notice

import (
	httpTypes "entdemo/common/types"
	serviceCtx "entdemo/svc"
)

func (n *Notice) Add(svc *serviceCtx.ServiceContext, req httpTypes.Request) (httpTypes.Response, error) {
	return httpTypes.Response{
		Code: httpTypes.EntSuccess,
		Msg:  httpTypes.CodeInfo[httpTypes.EntSuccess],
		Data: nil,
	}, nil
}
