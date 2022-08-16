package middleware

import (
	"net/http"
	"strconv"

	"ark-zero-admin/common/errorx"
	"ark-zero-admin/common/globalkey"
	"ark-zero-admin/common/utils"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type PermMenuAuthMiddleware struct {
	JwtSecret string
	Redis     *redis.Redis
}

func NewPermMenuAuthMiddleware(s string, r *redis.Redis) *PermMenuAuthMiddleware {
	return &PermMenuAuthMiddleware{
		JwtSecret: s,
		Redis:     r,
	}
}

func (m *PermMenuAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if len(r.Header.Get("Authorization")) > 0 {
			userId := utils.GetUserId(r.Context())
			is, err := m.Redis.Sismember(globalkey.CachePermMenuKey+strconv.FormatInt(userId, 10), r.RequestURI)
			if err != nil || is != true {
				httpx.Error(w, errorx.NewDefaultError(errorx.NotPermMenuErrorCode))
			} else {
				next(w, r)
			}
		}
	}
}
