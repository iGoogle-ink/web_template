package service

import (
	"runtime"
	"sync"

	"gotil/web"
	"web_template/config"
	"web_template/dao"
	"web_template/ecode"
	"web_template/model"

	"github.com/gin-gonic/gin"
)

type Service struct {
	rwMu    sync.RWMutex
	c       *config.Config
	dao     *dao.Dao
	appKeys map[string]struct{} // key:appkey
}

func New(c *config.Config) (svc *Service) {
	if c.NumCPU > 0 {
		runtime.GOMAXPROCS(c.NumCPU)
	}
	svc = &Service{
		c:       c,
		dao:     dao.New(c),
		appKeys: make(map[string]struct{}),
	}
	svc.setAppKeys()
	//s.loadErrorNumber()
	//go s.loadErrorNumberProc(1)
	return svc
}

func (s *Service) Close() {
	if s.dao != nil {
		s.dao.Close()
	}
}

// VerifyAppKey 验证 appkey
func (s *Service) VerifyAppKey(c *gin.Context) {
	appkey := c.GetHeader(model.HeaderAppKey)
	_, ok := s.appKeys[appkey]
	if !ok {
		web.JSON(c, nil, ecode.InvalidAppKeyErr)
		c.Abort()
		return
	}
}

func (s *Service) setAppKeys() {
	for _, v := range s.c.AppKey {
		s.appKeys[v] = struct{}{}
	}
}
