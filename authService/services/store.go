package services

import (
	"github.com/gin-gonic/contrib/sessions"
	"log"
)

var (
	_SessionStore sessions.Store
)

func SessionStore() sessions.Store {
	if _SessionStore == nil {
		s, err := sessions.NewRedisStore(5,"tcp","redis.default.svc.cluster.local:6379","",[]byte("auth"))
		if err != nil {
			log.Panicln(err)
		}
		_SessionStore = s
	}
	return _SessionStore
}
