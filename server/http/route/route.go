package route

import (
	httpTypes "entdemo/common/types" // import the types package
	"entdemo/internal/api/notice"    // import the notice package
	"entdemo/svc"                    // import the svc package
)

type ApiRoute struct {
	ApiGroups []ApiGroup
}

type ApiGroup struct {
	ServiceContext *svc.ServiceContext
	Method         string
	Path           string
	Handler        httpTypes.HandleFunc
}

func (r *ApiRoute) InitRoutes(s *svc.ServiceContext) error { // update the function signature
	notice, s := notice.NewNotice(s)
	r.AddRoute(s, "GET", "/api", notice.Add)
	return nil
}

func (r *ApiRoute) AddRoute(s *svc.ServiceContext, method string, path string, handler httpTypes.HandleFunc) {
	r.ApiGroups = append(r.ApiGroups, ApiGroup{
		ServiceContext: s,
		Method:         method,
		Path:           path,
		Handler:        handler,
	})
}
