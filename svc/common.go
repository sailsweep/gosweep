package svc

import (
	"context"
	myconfig "entdemo/config/types"
	svcTypes "entdemo/svc/types"
	"errors"
	"fmt"
)

type ServiceContext struct {
	Config        myconfig.Config
	Db            *svcTypes.Db
	Cache         *svcTypes.Cache
	internalClose []func()
	err           error
}

func NewServiceContext(ctx context.Context) (*ServiceContext, error) {
	var (
		service *ServiceContext
		c       myconfig.Config
		ok      bool
	)
	service = new(ServiceContext)
	c, ok = ctx.Value("config").(myconfig.Config)
	if !ok {
		service.err = errors.New("config not found")
		return service, service.err
	}
	service.Config = c
	return service, nil
}

func (s *ServiceContext) InitAllService(ctx context.Context) error {
	var (
		c myconfig.Config
	)
	if s.err != nil {
		return s.err
	}
	s.InitDb(c).InitCache(c)
	return s.err
}

func (s *ServiceContext) InitDb(c myconfig.Config) *ServiceContext {
	var (
		db       *svcTypes.Db
		dbConfig svcTypes.DbConfig
		err      error
	)
	if s.err != nil {
		return s
	}
	db = new(svcTypes.Db)
	dbConfig.Host = c.Database.Host
	dbConfig.Port = c.Database.Port
	dbConfig.DB = c.Database.DB
	dbConfig.User = c.Database.User
	dbConfig.Pass = c.Database.Pass
	db.Config = dbConfig
	err = db.Init()
	if err != nil {
		s.err = err
		return s
	}
	s.Db = db
	s.RegisterClose(s.Db)
	return s
}

func (s *ServiceContext) InitCache(c myconfig.Config) *ServiceContext {
	var (
		err         error
		cache       *svcTypes.Cache
		cacheConfig svcTypes.CacheConfig
	)
	if s.err != nil {
		return s
	}
	cache = new(svcTypes.Cache)
	cacheConfig.Host = c.Redis.Host
	cacheConfig.Port = c.Redis.Port
	cache.Config = cacheConfig
	err = cache.Init()
	if err != nil {
		s.err = err
		return s
	}
	s.Cache = cache
	s.RegisterClose(s.Cache)
	return s
}

func (s *ServiceContext) RegisterClose(in svcTypes.Service) error {
	closeFunc, err := in.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("register close func failed: %v", in))
	}
	s.internalClose = append(s.internalClose, closeFunc)
	return nil
}

func (s *ServiceContext) Close() error {
	for _, v := range s.internalClose {
		v()
	}
	return nil
}
