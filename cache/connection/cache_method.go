package connection

import (
	"github.com/teq-quocbang/store/cache"
	"github.com/teq-quocbang/store/cache/services/register"
)

func (r redisDB) Register() cache.RegisterService {
	return register.NewRedisRegisterCache(r.redis)
}
