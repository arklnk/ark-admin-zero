package middleware

import (
	"net/http"
	"strconv"

	"ark-admin-zero/common/errorx"
	"ark-admin-zero/common/globalkey"
	"ark-admin-zero/common/utils"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type PermMenuAuthMiddleware struct {
	Redis     *redis.Redis
}

func NewPermMenuAuthMiddleware(r *redis.Redis) *PermMenuAuthMiddleware {
	return &PermMenuAuthMiddleware{
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
