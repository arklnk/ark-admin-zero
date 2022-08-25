package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"ark-admin-zero/common/config"
	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/utils"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type PermMenuAuthMiddleware struct {
	Redis *redis.Redis
}

func NewPermMenuAuthMiddleware(r *redis.Redis) *PermMenuAuthMiddleware {
	return &PermMenuAuthMiddleware{
		Redis: r,
	}
}

func (m *PermMenuAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {
			userId := utils.GetUserId(r.Context())
			online, err := m.Redis.Get(config.SysOnlineUserCachePrefix + strconv.FormatUint(userId, 10))
			if err != nil || online == "" {
				httpx.Error(w, errorx.NewDefaultError(errorx.AuthErrorCode))
				var erring any
				erring = "Auth fail"
				panic(erring)
			}

			uri := strings.Split(r.RequestURI, "?")
			is, err := m.Redis.Sismember(config.SysPermMenuCachePrefix+strconv.FormatUint(userId, 10), uri[0])
			if err != nil || is != true {
				httpx.Error(w, errorx.NewDefaultError(errorx.NotPermMenuErrorCode))
			} else {
				next(w, r)
			}
		}
	}
}
